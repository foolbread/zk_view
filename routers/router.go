package routers

import (
	"zk_view/controllers"

	"path"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

func init() {
	beego.Router("/", controllers.GlobalController)
	beego.Router("/login", controllers.GlobalController, "post:Login")
	beego.Router("/path", controllers.GlobalController, "post:Path")
	beego.Router("/home", controllers.GlobalController, "get:Home")

	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/dist", "static/dist")

	beego.SetViewsPath("static")

	beego.AddFuncMap("GetDir", path.Dir)

	var conf session.ManagerConfig
	conf.CookieName = "gosessionid"
	conf.EnableSetCookie = true
	conf.Gclifetime = 3600
	conf.Maxlifetime = 3600
	conf.Secure = false
	conf.CookieLifeTime = 3600
	controllers.GlobalSession, _ = session.NewManager("memory", &conf)
	go controllers.GlobalSession.GC()
}
