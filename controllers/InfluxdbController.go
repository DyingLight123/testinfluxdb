package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/redis.v4"
	"log"
	"testinfluxdb/models"
	"time"
)

type InfluxdbController struct {
	beego.Controller
}

type InputTime struct {
	T1   string            `json:"t1"`
	T2   string            `json:"t2"`
	Tags map[string]string `json:"tags"`
}

func (maps *InfluxdbController) Get() {
	t := time.Now()
	err := models.AddInfluxdbData()
	if err == redis.Nil {
		maps.Data["json"] = map[string]interface{}{"message": "没有从redis查询到数据！"}
	} else if err == nil {
		maps.Data["json"] = map[string]interface{}{"message": "成功！"}
	} else {
		log.Fatal(err)
	}
	tt := time.Since(t)
	fmt.Println("总的时间：", tt)
	maps.ServeJSON()
}

func (maps *InfluxdbController) Post() {
	t := new(InputTime)
	json.Unmarshal(maps.Ctx.Input.RequestBody, t)
	result, err := models.GetInfluxdbData(t.T1, t.T2, t.Tags)
	if err != nil {
		log.Fatal(err)
	}
	list := make([]map[string]interface{}, 0)
	for i := 0; i < len(result); i++ {
		m := make(map[string]interface{})
		j, _ := json.Marshal(&result[i])
		json.Unmarshal(j, &m)
		list = append(list, m)
	}
	maps.Data["json"] = map[string]interface{}{"data": list, "length": len(list)}
	maps.ServeJSON()
}
