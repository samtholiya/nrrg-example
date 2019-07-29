package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

//Redis struct holding the connection obj for redis cache
type Redis struct {
	client *redis.Client
}

//SetConfig sets the client configuration
func (r *Redis) SetConfig(address, password string) {
	r.client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

//Set sets the string value in redis
func (r *Redis) Set(key, value string) error {
	fmt.Println(key, value)
	return r.client.Set(key, value, 0).Err()
}

//Get gets the string corresponding to the key provided
func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}
