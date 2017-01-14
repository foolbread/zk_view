package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	IsLogin bool
	beego.Controller
}

func (c *BaseController) Prepare(){
	userLogin := c.GetSession("userLogin")
	if userLogin == nil{
		c.IsLogin = false
	}else {
		c.IsLogin = true
		c.Data["UserName"] = userLogin
	}
}