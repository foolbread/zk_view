package controllers

import (
	"github.com/astaxie/beego/session"
)

var GlobalSession *session.Manager

var GlobalController *MainController = new(MainController)
