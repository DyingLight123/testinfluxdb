package models

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

func AddRedisData(number int) error {
	cli := ConnRedis()
	defer cli.Close()

	_, err := cli.Del("data").Result()
	if err != nil {
		return err
	}
	rand.Seed(time.Now().Unix())
	for i := 0; i < number; i++ {
		x := rand.Float64() * 1000000
		m := make(map[string]interface{})
		m["value"+strconv.Itoa(i)] = x
		j, _ := json.Marshal(m)
		cli.HSet("data", strconv.Itoa(i), string(j))
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
