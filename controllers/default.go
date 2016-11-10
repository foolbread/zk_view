package controllers

import (
	"fmt"

	"github.com/foolbread/zk_view/models"

	"path"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	sid := c.Ctx.GetCookie("gosessionid")
	if len(sid) == 0 {
		c.TplName = "login.html"
		return
	}

	fmt.Println("SID:", sid)

	sess, _ := GlobalSession.GetSessionStore(sid)
	name := sess.Get("username")
	if name == nil {
		c.TplName = "login.html"
		return
	}

	c.Data["UserName"] = name

}

func (c *MainController) Get() {
	if len(c.TplName) > 0 {
		return
	}

	c.Redirect("/home?zkPath=/", 302)
}

func (c *MainController) Path() {
	action := c.GetString("action")
	curPath := c.GetString("currentPath")

	switch action {
	case "add":
		ap := c.GetString("addPath")
		av := c.GetString("addValue")
		str, err := models.CreatePath(path.Join(curPath, ap), av)
		fmt.Println(path.Join(curPath, ap))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(str)
	case "update":
		up := c.GetString("updatePath")
		uv := c.GetString("updateValue")
		var err error
		if up == "[.]" {
			err = models.UpdatePathVal(curPath, uv)
		} else {
			err = models.UpdatePathVal(path.Join(curPath, up), uv)
		}
		if err != nil {
			fmt.Println(err)
		}
	case "delete":
		dps := c.GetStrings("nodeChkGroup")
		for _, v := range dps {
			models.DeletePath(v)
		}
	}

	c.Redirect(fmt.Sprintf("/home?zkPath=%s", curPath), 302)
}

func (c *MainController) Home() {
	if len(c.TplName) > 0 {
		return
	}

	fmt.Println("zkpath:", c.Ctx.Request.URL.Query().Get("zkPath"))

	pa := c.Ctx.Request.URL.Query().Get("zkPath")
	subpaths, zkPairs, err := models.GetPathInfo(pa)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Data["CurZKPath"] = pa
	c.Data["SubZKPaths"] = subpaths
	c.Data["CurPath"] = zkPairs[0].Path
	c.Data["CurPathVal"] = zkPairs[0].Val
	c.Data["SubZKPathVals"] = zkPairs[1:]
	c.TplName = "home.html"
}

func (c *MainController) Login() {
	user := c.GetString("user")
	pwd := c.GetString("pwd")
	if !models.GetUserInstance().CheckUser(user, pwd) {
		c.Ctx.WriteString("user or password is error!")
		return
	}

	sess, _ := GlobalSession.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	sess.Set("username", user)

	fmt.Println("name:", user, "pwd:", pwd)

	c.Data["UserName"] = user
	c.Redirect("/home?zkPath=/", 302)
}
