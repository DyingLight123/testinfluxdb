package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/robfig/cron"
	"gopkg.in/redis.v4"
	"log"
	"testinfluxdb/models"
	"time"
)

type RefreshController struct {
	beego.Controller
}

var pause1 = make(chan string, 1)
var status1 int

func (data *RefreshController) Get() {
	conn1 := models.ConnInfluxdb()
	conn2 := models.ConnRedis()
	go RefreshInfluxdb(conn1, conn2)

	data.Data["json"] = map[string]interface{}{}
	data.ServeJSON()
}

func (data *RefreshController) Post() {
	err := PauseInfluxdb()
	if err != nil {
		data.Data["json"] = map[string]interface{}{"message": "可以开始刷新了！"}
	} else {
		data.Data["json"] = map[string]interface{}{"message": "刷新停止！"}
	}
	data.ServeJSON()
}

func RefreshInfluxdb(conn1 client.Client, conn2 *redis.Client) {
	if status1 == 1 {
		log.Println("refreshing! please pause! ")
		return
	}
	status1 = 1
	c := cron.New()
	c.AddFunc("@every "+"60s", func() {
		t := time.Now()
		err := models.AddInfluxdbData(conn1, conn2)
		if err != nil {
			log.Println(err)
			return
		}
		tt := time.Since(t)
		log.Println("总的时间：", tt)
		log.Println("")
	})
	c.Start()
	<-pause1
	status1 = 0
	log.Println("continue")
	c.Stop()
	//time.AfterFunc(30 * time.Second, c.Stop)
}

func PauseInfluxdb() error {
	if status1 != 1 {
		return errors.New("please begin refresh! ")
	}
	log.Println("pause")
	pause1 <- "continue"
	return nil
}
