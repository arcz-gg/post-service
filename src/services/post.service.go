package postservices

import (
	"context"

	"github.com/Arcz.gg/post-service/proto"
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
	post_proto.UnimplementedPostServiceServer
}

func (s *PostService) CreatePost(
	ctx context.Context, 
	req *post_proto.CreatePostRequest,
) (*post_proto.CreatePostResponse, error) {
	
	data := req.GetPost(); 

	id := data.GetUserId() + "-evt-" + utils.GenId(6);
	res := &post_proto.CreatePostResponse{

		PostId: id,
		Created: true,
	}

	return res, nil;
}
