package models

import (
	"strings"

	"github.com/foolbread/fbcommon/golog"

	"encoding/json"

	"github.com/astaxie/beego"
)

func InitModels() {
	str := beego.AppConfig.String("zkserver")
	g_zk = newZKCon(strings.Split(str, ";"))
	err := g_zk.connectZK()
	if err != nil {
		golog.Critical(err)
	}

	str = beego.AppConfig.String("auth")
	err = json.Unmarshal([]byte(str), &g_userManager)
	if err != nil {
		golog.Critical(err)
	}

	for _, v := range g_userManager.Users {
		golog.Info("user:", v.User, "pwd:", v.Pwd)
	}
}

func GetZKInstance() *zkCon {
	return g_zk
}

func GetUserInstance() *webUserManager {
	return &g_userManager
}

var g_zk *zkCon

var g_userManager webUserManager
