package models

import (
	"path"
	"sync"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var zk_server []string = []string{"192.168.250.178:2181", "192.168.250.195:2181"}

type ZKPair struct {
	Path string
	Val  string
}

type zkCon struct {
	lo    *sync.Mutex
	con   *zk.Conn
	addrs []string
}

func newZKCon(addrs []string) *zkCon {
	r := new(zkCon)
	r.lo = new(sync.Mutex)
	r.addrs = addrs

	return r
}

func (z *zkCon) connectZK() error {
	con, _, err := zk.Connect(z.addrs, 20*time.Second)
	if err != nil {
		return err
	}

	z.con = con
	return nil
}

func (z *zkCon) reconnectZK() error {
	z.lo.Lock()
	defer z.lo.Unlock()

	z.closeZK()

	return z.connectZK()
}

func (z *zkCon) closeZK() {
	if z.con != nil {
		z.con.Close()
	}
}

func (z *zkCon) getVal(path string) ([]byte, *zk.Stat, error) {
	z.lo.Lock()
	data, stat, err := z.con.Get(path)
	z.lo.Unlock()
	if err != nil {
		return nil, nil, err
	}
	return data, stat, nil
}

func (z *zkCon) getChilds(path string) ([]string, *zk.Stat, error) {
	z.lo.Lock()
	data, stat, err := z.con.Children(path)
	z.lo.Unlock()
	if err != nil {
		return nil, nil, err
	}

	return data, stat, nil
}

func (z *zkCon) updateVal(pa string, val string) (*zk.Stat, error) {
	z.lo.Lock()
	stat, err := z.con.Set(pa, []byte(val), -1)
	z.lo.Unlock()
	if err != nil {
		return nil, err
	}

	return stat, nil
}

func (z *zkCon) createPath(pa string, val string) (string, error) {
	return z.con.Create(pa, []byte(val), 0, zk.WorldACL(zk.PermAll))
}

func (z *zkCon) deletePath(pa string) error {
	z.lo.Lock()
	err := z.con.Delete(pa, -1)
	z.lo.Unlock()

	return err
}

//////////////////////////////////////////////////////////////////////////////////////////////////
func CreatePath(pa string, val string) (string, error) {
	return g_zk.createPath(pa, val)
}

func DeletePath(pa string) error {
	childs, _, err := g_zk.getChilds(pa)
	if err != nil {
		return err
	}

	if len(childs) == 0 {
		return g_zk.deletePath(pa)
	}

	for _, v := range childs {
		err = DeletePath(path.Join(pa, v))
		if err != nil {
			return err
		}
	}

	return g_zk.deletePath(pa)
}

func UpdatePathVal(pa string, val string) error {
	_, err := g_zk.updateVal(pa, val)
	return err
}

func GetPathInfo(pa string) ([]string, []*ZKPair, error) {
	childs, _, err := g_zk.getChilds(pa)
	if err != nil {
		return nil, nil, err
	}

	var retPairs []*ZKPair
	data, _, err := g_zk.getVal(pa)
	if err != nil {
		return nil, nil, err
	}

	retPairs = append(retPairs, &ZKPair{path.Base(pa), string(data)})

	for _, v := range childs {
		data, _, err := g_zk.getVal(path.Join(pa, v))
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
