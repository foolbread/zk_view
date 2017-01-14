package user

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/foolbread/fbcommon/golog"
)

func InitUser(){
	g_user = new(webUserManager)

	str := beego.AppConfig.String("auth")
	err := json.Unmarshal([]byte(str), g_user)
	if err != nil {
		golog.Critical(err)
	}

	for _, v := range g_user.Users {
		golog.Info("user:", v.User, "pwd:", v.Pwd)
	}
}

type webUser struct {
	User string `json:"user"`
	Pwd  string `json:"pwd"`
}

type webUserManager struct {
	Users []*webUser `json:"users"`
}

func GetUserManager()*webUserManager{
	return g_user
}

var g_user *webUserManager

func (u *webUserManager)CheckLogin(usr string, pwd string)bool{
	for _, v := range u.Users {
		if v.User == usr {
			return v.Pwd == pwd
		}
	}

	return false
}
