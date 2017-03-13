package models

import (
	. "zk_view/models/user"
	. "zk_view/models/zookeeper"
	. "zk_view/models/watch"
)

func InitModels() {
	InitUser()
	InitZookeeper()
	InitWatch()
}

