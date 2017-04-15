package user

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/foolbread/fbcommon/golog"
	"strings"
)

func InitUser(){
	g_user = new(webUserManager)
	g_user.UserMap = make(map[string]*webUser)

	str := beego.AppConfig.String("auth")
	err := json.Unmarshal([]byte(str), g_user)
	if err != nil {
		golog.Critical(err)
	}

	for _, v := range g_user.Users {
		golog.Info("user:", v.User, "pwd:", v.Pwd,"path:",v.Path)
		g_user.UserMap[v.User] = &webUser{v.User,v.Pwd,v.Path}
	}
}

type webUser struct {
	User string `json:"user"`
	Pwd  string `json:"pwd"`
	Path string `json:"path"`
}

type webUserManager struct {
	Users []*webUser `json:"users"`
	UserMap map[string]*webUser
}

func GetUserManager()*webUserManager{
	return g_user
}

var g_user *webUserManager

func (u *webUserManager)CheckLogin(usr string, pwd string)bool{
	info := u.getUserInfo(usr)
	if info == nil{
		return false
	}

	return info.Pwd == pwd
}

func (u *webUserManager)CheckPath(usr string, pa string)bool{
	info := u.getUserInfo(usr)
	if info == nil{
		return false
	}

	return strings.Contains(pa,info.Path)
}

func (u *webUserManager)GetUsrBasePath(usr string)string{
	info := u.getUserInfo(usr)
	if info == nil{
		return ""
	}

	return info.Path
}

func (u *webUserManager)getUserInfo(usr string)*webUser{
	return u.UserMap[usr]
}