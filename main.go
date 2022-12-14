package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/Arcz.gg/post-service/certifs"
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

	creds, err := credentials.NewServerTLSFromFile(certifs.Path("./server_cert.pem"), certifs.Path("./server_key.pem"));
	isError(err);

	opts := []grpc.ServerOption{
		grpc.MaxConcurrentStreams(100),
		grpc.Creds(creds),
	};

	server := grpc.NewServer(opts...);
	post_proto.RegisterPostServiceServer(server, &postservices.PostServiceDescriptor{});
	
	log.Printf("[SERVER]: Server started at port localhost%s", port)
	if err := server.Serve(ln); err != nil {
		log.Fatalf("[ERROR]: Failed to serve %s", err)
	}
}
