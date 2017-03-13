package main

import (
	"github.com/foolbread/zk_view/models"
	_ "github.com/foolbread/zk_view/routers"
	_ "github.com/foolbread/zk_view/initial"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	models.InitModels()

	beego.InsertFilter("/*",beego.BeforeRouter, FilterUser)

	beego.Run()
}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("userLogin").(string)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}