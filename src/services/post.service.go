package postservices

import (
	"context"

	_ "google.golang.org/grpc"
	"github.com/Arcz.gg/post-service/proto"
)

type PostServiceDescriptor struct {
	post_proto.UnimplementedPostServiceServer
}

type Post struct {
	ID string;
	UserId string;
	data string;
	DateCreated string;
};

type CreatePostRequest struct {
	
	UserId string;
	data string;
}

type CreatePostResponse struct {

	PostId string
}

type PostService struct {}

func (s * PostService) CreatePost(
	ctx *context.Context, 
	req *CreatePostRequest,
) (* CreatePostResponse, error) {

	return &CreatePostResponse{
		PostId: "12345",
	}, nil
}
