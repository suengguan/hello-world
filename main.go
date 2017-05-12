package main

import (
	_ "hello-world/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Debug("this is hello world-v2")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
