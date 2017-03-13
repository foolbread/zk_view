package initial

import (
	"github.com/astaxie/beego"
)

func InitStatic(){
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/dist", "static/dist")

	beego.SetViewsPath("static")
}
