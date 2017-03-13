package zookeeper

import (
	"github.com/foolbread/fbcommon/golog"

	"zk_view/controllers"
	. "zk_view/models/zookeeper"

	"path"
)

type ZooKeeperController struct {
	controllers.BaseController
}

func (z *ZooKeeperController)Get(){
	golog.Info("zkpath:", z.Ctx.Request.URL.Query().Get("zkPath"))

	pa := z.Ctx.Request.URL.Query().Get("zkPath")
	subpaths, zkPairs, err := GetPathInfo(pa)
	if err != nil {
		golog.Error(err)
		return
	}

	var hrefPaths []string
	var tmpPaths []string
	var dir string = pa
	for {
		tmpPaths = append(tmpPaths, dir)
		if dir == "/" {
			for i := len(tmpPaths) - 1; i >= 0; i-- {
				hrefPaths = append(hrefPaths, tmpPaths[i])
			}
			break
		}

		dir = path.Dir(dir)
	}

	z.Data["HrefPaths"] = hrefPaths
	z.Data["CurZKPath"] = pa
	z.Data["SubZKPaths"] = subpaths
	z.Data["CurPath"] = zkPairs[0].Path
	z.Data["CurPathVal"] = zkPairs[0].Val
	z.Data["SubZKPathVals"] = zkPairs[1:]
	z.TplName = "home.html"
}
