package zookeeper

import (
	"github.com/astaxie/beego"
	"strings"
	"path"
	"sort"

	"zk_view/utils"
	"github.com/foolbread/fbcommon/golog"
)

func InitZookeeper(){
	str := beego.AppConfig.String("zkserver")
	addrs := strings.Split(str, ",")

	g_zk = utils.NewZKCon(addrs)
	err := g_zk.ConnectZK()
	if err != nil{
		golog.Critical(err)
	}
}

type ZKPair struct {
	Path string
	Val  string
}

var g_zk *utils.ZKCon

func CreatePath(pa string, val string) (string, error) {
	return g_zk.CreatePath(pa, val)
}

func DeletePath(pa string) error {
	childs, _, err := g_zk.GetChilds(pa)
	if err != nil {
		return err
	}

	if len(childs) == 0 {
		return g_zk.DeletePath(pa)
	}

	for _, v := range childs {
		err = DeletePath(path.Join(pa, v))
		if err != nil {
			return err
		}
	}

	return g_zk.DeletePath(pa)
}

func UpdatePathVal(pa string, val string) error {
	_, err := g_zk.UpdateVal(pa, val)
	return err
}

func GetPathInfo(pa string) ([]string, []*ZKPair, error) {
	childs, _, err := g_zk.GetChilds(pa)
	if err != nil {
		return nil, nil, err
	}

	var retPairs []*ZKPair
	data, _, err := g_zk.GetVal(pa)
	if err != nil {
		return nil, nil, err
	}

	retPairs = append(retPairs, &ZKPair{path.Base(pa), string(data)})

	sort.Strings(childs)
	for _, v := range childs {
		data, _, err := g_zk.GetVal(path.Join(pa, v))
		if err != nil {
			return nil, nil, err
		}

		p := new(ZKPair)
		p.Path = v
		p.Val = string(data)
		retPairs = append(retPairs, p)
	}

	return childs, retPairs, nil
}