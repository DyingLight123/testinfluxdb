package controllers

import (
	"github.com/DyingLight123/RedisInfluxdb"
	"github.com/astaxie/beego"
)

type OverController struct {
	beego.Controller
	Redisinfluxdb *RedisInfluxdb.RedisInfluxdb
}

func (this *OverController) Post() {
	err := this.Redisinfluxdb.PauseRedis()
	if err != nil {
		this.Data["json"] = map[string]interface{}{"message": "请先开始刷新！"}
	} else {
		this.Data["json"] = map[string]interface{}{"message": "刷新停止！"}
	}
	this.ServeJSON()
}
