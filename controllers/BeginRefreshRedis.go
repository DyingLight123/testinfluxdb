package controllers

import (
	"fmt"
	"github.com/DyingLight123/RedisInfluxdb"
	"github.com/astaxie/beego"
)

type BeginRefreshRedis struct {
	beego.Controller
	Redisinfluxdb *RedisInfluxdb.RedisInfluxdb
}

/*func (this *BeginRefreshRedis) Prepare() {
	this.Redisinfluxdb = &RedisInfluxdb.RedisInfluxdb{"127.0.0.1:6379", "", "data",
		"http://127.0.0.1:8086", "admin", "admin", "test",
		"redis"}
	fmt.Println(this.Redisinfluxdb)
}*/

func (this *BeginRefreshRedis) Get() {
	this.Redisinfluxdb = &RedisInfluxdb.RedisInfluxdb{"localhost:6379", "", "data",
		"http://127.0.0.1:8086", "admin", "admin", "test",
		"redis"}
	fmt.Println(this.Redisinfluxdb)

	go this.Redisinfluxdb.RefreshRedis(100000)

	this.Data["json"] = map[string]interface{}{}
	this.ServeJSON()
}

func (this *BeginRefreshRedis) Post() {
	err := this.Redisinfluxdb.PauseRedis()
	if err != nil {
		this.Data["json"] = map[string]interface{}{"message": "可以开始刷新了！"}
	} else {
		this.Data["json"] = map[string]interface{}{"message": "刷新停止！"}
	}
	this.ServeJSON()
}
