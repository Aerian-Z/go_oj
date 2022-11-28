package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func TestRedisSet(t *testing.T) {
	err := rdb.Set(ctx, "key", "value", 10*time.Second).Err()
	if err != nil {
		t.Error(err)
	}
}

func TestRedisGet(t *testing.T) {
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("key", val)
}
