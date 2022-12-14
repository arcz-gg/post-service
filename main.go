package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/Arcz.gg/post-service/proto"
	"github.com/Arcz.gg/post-service/src/services"
	"github.com/Arcz.gg/post-service/src/store"
	cache "github.com/go-redis/cache/v9"
	"github.com/go-redis/redis/v9"
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
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6373",
	});


	defer rdb.Close();

	storeCache := cache.New(&cache.Options{Redis: rdb});
	stores := store.Store{Rdb: rdb, CStore: storeCache};
	service := &postservices.PostService{
		Stores: &stores,
	};

	post_proto.RegisterPostServiceServer(server, service);	
	log.Printf("[SERVER]: Server started at port localhost%s", port)
	if err := server.Serve(ln); err != nil {
		log.Fatalf("[ERROR]: Failed to serve %s", err)
	}
}
