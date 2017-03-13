package routers

import (
	"zk_view/controllers/user"
	"zk_view/controllers/zookeeper"
	"zk_view/controllers/watch"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &user.LoginUserController{})
	beego.Router("/logout", &user.LogoutUserController{})
	beego.Router("/home", &zookeeper.ZooKeeperController{})
	beego.Router("/path", &zookeeper.ZooKeeperPathController{})
	beego.Router("/check", &watch.WatchCheckController{})
}
