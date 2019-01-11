package routers

import (
	"github.com/astaxie/beego"
	"testinfluxdb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/influxdb", &controllers.InfluxdbController{})

	beego.Router("/refresh", &controllers.RedisController{})

	//beego.Router("/refresh", &controllers.BeginRefreshRedis{})
}
