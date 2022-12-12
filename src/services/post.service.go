package postservices

import (
	"context"
	"log"

	"github.com/Arcz.gg/post-service/proto"
	"github.com/google/uuid"
	_ "google.golang.org/grpc"
)

type Post struct {
	ID string;
	UserId string;
	data string;
	DateCreated string;
};

type CreatePostRequest struct {

	Post *Post;
}

type CreatePostResponse struct {

	PostId string;
	Created bool;
}

type PostService struct {
	post_proto.UnimplementedPostServiceServer
}

func (s *PostService) CreatePost(
	ctx context.Context, 
	req *post_proto.CreatePostRequest,
) (*post_proto.CreatePostResponse, error) {
	
	data := req.GetPost();
	log.Println("[LOG]: ", data);

	id, err := uuid.NewRandom();
	if err != nil {

		return nil, err;
	}

	log.Println("[LOG]: New UUID Generated ", id);

	res := &post_proto.CreatePostResponse{
		PostId: data.Id,
		Created: true,
	}

	return res, nil;
}
