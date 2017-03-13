package user

import (
	"github.com/foolbread/zk_view/controllers"

	. "github.com/foolbread/zk_view/models/user"
)

//登录
type LoginUserController struct {
	controllers.BaseController
}

func (u *LoginUserController) Get(){
	check := u.BaseController.IsLogin
	if check {
		u.Redirect("/home?zkPath=/", 302)
	}else{
		u.TplName = "login.html"
	}
}

func (u *LoginUserController) Post(){
	username := u.GetString("user")
	passwd := u.GetString("pwd")

	ok := GetUserManager().CheckLogin(username, passwd)
	if ok{
		u.SetSession("userLogin", username)
		u.Redirect("/home?zkPath=/", 302)
	}else{
		u.Data["message"] = "login failed!"
		u.TplName = "login.html"
	}
}


//退出
type LogoutUserController struct {
	controllers.BaseController
}

func (u *LogoutUserController)Get(){
	u.DelSession("userLogin")
	u.Redirect("/login",302)
}