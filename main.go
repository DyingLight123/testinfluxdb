package main

import (
	"github.com/astaxie/beego"
	_ "testinfluxdb/routers"
)

func main() {
	beego.Run()
}

