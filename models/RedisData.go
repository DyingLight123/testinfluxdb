package models

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

func AddRedisData(number int) error {
	redis := ConnRedis()
	defer redis.Close()
	_, err := redis.Del("data").Result()
	if err != nil {
		return err
	}
	rand.Seed(time.Now().Unix())
	for i := 0; i < number; i++ {
		x := rand.Intn(10000)
		m := make(map[string]interface{})
		m["value"] = x
		j, _ := json.Marshal(m)
		redis.HSet("data", "value" + strconv.Itoa(i), string(j))
	}
	return nil
}

func GetRedisData() (map[string]string, error) {
	client := ConnRedis()
	defer client.Close()
	/*field, err := client.HGet("map", "value99").Result()
	if err == redis.Nil  {
		return nil, nil
	} else if err == nil {
		m := make(map[string]interface{})
		json.Unmarshal([]byte(field), &m)
		return m, nil
	} else {
		return nil, err
	}*/
	field, err := client.HGetAll("data").Result()
	return field, err
}