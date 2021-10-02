package infrastructure

import (
	"github.com/gin-contrib/sessions/redis"
)

func InitRedis(host string) redis.Store {
	store, err := redis.NewStore(10, "tcp", host, "", []byte("secret"))
	if err != nil {
		panic("Fail to connect Cache.")
	}

	return store
}
