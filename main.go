package main

import (
	_ "opcode_server/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Debug("this is hello world-v57")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
