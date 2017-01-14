package initial

import (
	"path"
	"github.com/astaxie/beego"
)

func InitTemplate(){
	beego.AddFuncMap("GetDir", path.Dir)
	beego.AddFuncMap("GetBase", path.Base)
}
