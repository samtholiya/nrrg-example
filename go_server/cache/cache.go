package cache

import (
	"github.com/samtholiya/tempServer/cache/redis"
)

//Database every cache database should implement
type Database interface {

	//SetConfig set the config of cache database
	SetConfig(address, password string)

	//Set sets the string value in redis
	Set(key, value string) error

	//Get gets the string corresponding to the key provided
	Get(key string) (string, error)
}

//New returns a cache database object
func New() Database {
	return &redis.Redis{}
}
