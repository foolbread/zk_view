package watch

import (
	"zk_view/controllers"

	. "zk_view/models/watch"
)

type WatchCheckController struct {
	controllers.BaseController
}

func (w *WatchCheckController)Get(){
	w.Ctx.WriteString(CheckWatch())
}
