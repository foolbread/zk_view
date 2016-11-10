package models

import (
	"strings"

	"github.com/foolbread/fbcommon/golog"

	"github.com/astaxie/beego"
)

func InitModels() {
	str := beego.AppConfig.String("zkserver")
	g_zk = newZKCon(strings.Split(str, ";"))
	err := g_zk.connectZK()
	if err != nil {
		golog.Critical(err)
	}
}

var g_zk *zkCon
