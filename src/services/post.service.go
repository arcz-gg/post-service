package postservices

import (
	"context"
	"log"

	"github.com/Arcz.gg/post-service/proto"
	"github.com/Arcz.gg/post-service/src/store"
	"github.com/Arcz.gg/post-service/src/utils"
	_ "google.golang.org/grpc"
)

type Post struct {
	ID string;
	UserId string;
	data string;
	DateCreated string;
};

type PostService struct {
	post_proto.UnimplementedPostServiceServer;
	Stores store.PostStorage;
}

func (s *PostService) CreatePost(
	ctx context.Context, 
	req *post_proto.CreatePostRequest,
) (*post_proto.CreatePostResponse, error) {
	
	data := req.GetPost(); 

	id := data.GetUserId() + "-evt-" + utils.GenId(6);
	data.Id = id;
	
	res := &post_proto.CreatePostResponse{

		PostId: id,
		Created: true,
	}
	err := s.Stores.Save(ctx, data, id);
	if err != nil {
		log.Println(err);
		return nil, err;
	}
	return res, nil;
}
