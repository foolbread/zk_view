package zookeeper

import (
	"fmt"
	"path"

	"zk_view/controllers"

	"github.com/foolbread/fbcommon/golog"

	. "zk_view/models/zookeeper"
)

type ZooKeeperPathController struct {
	controllers.BaseController
}

func (z *ZooKeeperPathController) Post() {
	action := z.GetString("action")
	curPath := z.GetString("currentPath")

	z.BaseController
	switch action {
	case "add":
		ap := z.GetString("addPath")
		av := z.GetString("addValue")
		str, err := CreatePath(path.Join(curPath, ap), av)
		golog.Info(path.Join(curPath, ap))
		if err != nil {
			golog.Error(err)
		}
		golog.Info(str)
	case "update":
		up := z.GetString("updatePath")
		uv := z.GetString("updateValue")
		var err error
		if up == "[.]" {
			err = UpdatePathVal(curPath, uv)
		} else {
			err = UpdatePathVal(path.Join(curPath, up), uv)
		}
		if err != nil {
			golog.Error(err)
		}
	case "delete":
		dps := z.GetStrings("nodeChkGroup")
		for _, v := range dps {
			DeletePath(v)
		}
	}

	z.Redirect(fmt.Sprintf("/home?zkPath=%s", curPath), 302)
}
