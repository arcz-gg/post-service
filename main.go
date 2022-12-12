package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	
	"github.com/Arcz.gg/post-service/proto"
	"github.com/Arcz.gg/post-service/src/services"
	)

func isError(err error) {
	
	if (err != nil) {
		log.Panic("[Error]: ", err);
	}
}
func main() {
	port := ":8080"
	ln, err := net.Listen("tcp", port);
	isError(err);
	
	server := grpc.NewServer(grpc.MaxConcurrentStreams(100));
	post_proto.RegisterPostServiceServer(server, &postservices.PostServiceDescriptor{});
	
	log.Printf("[SERVER]: Server started at port localhost%s", port)
	if err := server.Serve(ln); err != nil {
		log.Fatalf("[ERROR]: Failed to serve %s", err)
	}
}
