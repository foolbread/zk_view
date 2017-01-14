package routers

import (
	"github.com/foolbread/zk_view/controllers/user"
	"github.com/foolbread/zk_view/controllers/zookeeper"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &user.LoginUserController{})
	beego.Router("/logout", &user.LogoutUserController{})
	beego.Router("/home", &zookeeper.ZooKeeperController{})
	beego.Router("/path", &zookeeper.ZooKeeperPathController{})
}
