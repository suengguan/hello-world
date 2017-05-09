package main

import (
	_ "hello-world/routers"

	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("error message")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
