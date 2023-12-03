package databases

import (
	"log"

	"github.com/go-redis/redis/v8"

	"service2/configs"
	"service2/pkg/utils"

)

func NewRedisClient(cfg *configs.Config) *redis.Client {
	url, err := utils.UrlBuilder("redis", cfg)
	if err != nil {
		log.Fatal(err)
	}
	return redis.NewClient(&redis.Options{
		Addr: url,
	})
}
