package controllers

import (
	"github.com/astaxie/beego"
	"github.com/influxdata/platform/kit/errors"
	"github.com/robfig/cron"
	"gopkg.in/redis.v4"
	"log"
	"testinfluxdb/models"
)

type RedisController struct {
	beego.Controller
}

/*type InPutInt struct {
	Number  int    `json:"number"`
	Refresh string `json:"refresh"`
}*/

var pause = make(chan string, 1)
var status int

func (data *RedisController) Get() {
	conn2 := models.ConnRedis()
	go RefreshRedis(100000, conn2)

	data.Data["json"] = map[string]interface{}{}
	data.ServeJSON()
}

func (data *RedisController) Post() {
	err := PauseRedis()
	if err != nil {
		data.Data["json"] = map[string]interface{}{"message": "可以开始刷新了！"}
	} else {
		data.Data["json"] = map[string]interface{}{"message": "刷新停止！"}
	}
	data.ServeJSON()
}

func RefreshRedis(number int, conn2 *redis.Client) {
	if status == 1 {
		log.Println("refreshing! please pause! ")
		return
	}
	status = 1
	c := cron.New()
	c.AddFunc("@every "+"60s", func() {
		err := models.AddRedisData(number, conn2)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("redis已更新！")
	})
	c.Start()
	<-pause
	status = 0
	log.Println("continue")
	c.Stop()
	//time.AfterFunc(30 * time.Second, c.Stop)
}

func PauseRedis() error {
	if status != 1 {
		return errors.New("please begin refresh!")
	}
	log.Println("pause")
	pause <- "continue"
	return nil
}

/*func (data *RedisController) Get() {
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
*/
