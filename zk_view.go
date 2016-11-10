package main

import (
	"zk_view/models"
	_ "zk_view/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.InitModels()
	beego.Run()
}
