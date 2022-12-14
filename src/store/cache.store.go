package store

import (
	"context"

	post_proto "github.com/Arcz.gg/post-service/proto"
	cache "github.com/go-redis/cache/v9"
	"github.com/go-redis/redis/v9"
)

type PostStorage interface {

	Save(ctx context.Context, post *post_proto.Post, id string) error;
	Get(ctx context.Context, id string) ([]byte, error);
	Del(ctx context.Context, id string) error;
}

type Store struct {

	Rdb *redis.Client;
	CStore *cache.Cache;
}

func (s *Store) Save(ctx context.Context, post *post_proto.Post, id string) error { //will change to byte[]
	
	var dest = make([]byte, 0);
	item := cache.Item{
		Key: id,
		Value: dest,
		Do: func(i *cache.Item) (interface{}, error) {
			return s.CStore.Marshal(post);
		},
	}
	return s.CStore.Set(&item);
}

func (s *Store) Get(ctx context.Context, id string) ([]byte, error) {
	
	var target []byte;
	if err := s.CStore.Get(ctx, id, &target); err != nil {
		return nil, err;
	}

	return target, nil;
}

func (s *Store) Del(ctx context.Context, id string) error {

	return s.CStore.Delete(ctx, id);
}
