package routers

import (
	"github.com/astaxie/beego"
	"testinfluxdb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/influxdb", &controllers.InfluxdbController{})

	beego.Router("/redis", &controllers.RedisController{})

	beego.Router("/refresh", &controllers.RefreshController{})
}
