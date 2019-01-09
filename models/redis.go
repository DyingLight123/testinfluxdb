package models

import (
	"gopkg.in/redis.v4"
)

func ConnRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	return client
}
