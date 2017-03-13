package watch

import (
	"net/http"
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beego"
	"github.com/foolbread/fbcommon/golog"
)

func InitWatch(){
	g_watchurl = beego.AppConfig.String("watchurl")
	golog.Info("watchurl:", g_watchurl)
}

var g_watchurl string

func CheckWatch() string {
	res,err := http.Get(fmt.Sprintf("%s/%s", g_watchurl, "check"))
	if err != nil{
		return err.Error()
	}

	data,err := ioutil.ReadAll(res.Body)
	if err != nil{
		return err.Error()
	}

	res.Body.Close()
	return string(data)
}
