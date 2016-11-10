package main

import (
	"github.com/foolbread/zk_view/models"
	_ "github.com/foolbread/zk_view/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.InitModels()
	beego.Run()
}
