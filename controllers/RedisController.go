package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/robfig/cron"
	"log"
	"testinfluxdb/models"
	"time"
)

type RedisController struct {
	beego.Controller
}

type InPutInt struct {
	Number  int    `json:"number"`
	Refresh string `json:"refresh"`
}

func (data *RedisController) Post() {
	var request InPutInt
	json.Unmarshal(data.Ctx.Input.RequestBody, &request)
	c := cron.New()
	c.AddFunc("@every " + request.Refresh, func() {
		err := models.AddRedisData(request.Number)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("redis已更新！")
	})
	c.Start()
	time.AfterFunc(time.Minute, c.Stop)
	data.ServeJSON()
}

func (data *RedisController) Get() {
	field, err := models.GetRedisData()
	list := make(map[string]interface{}, 0)
	if err != nil {
		log.Fatal(err)
	} else if field == nil {
		data.Data["json"] = map[string]interface{}{"message": "没有数据！"}
	} else {
		for key, value := range field {
			m := make(map[string]interface{})
			json.Unmarshal([]byte(value), &m)
			list[key] = m
		}
		data.Data["json"] = list
	}
	//fmt.Println(field)
	data.ServeJSON()
}