package utils

import (
	"github.com/samuel/go-zookeeper/zk"
	"sync"
	"time"
)

type ZKCon struct {
	lo    *sync.Mutex
	con   *zk.Conn
	addrs []string
}

func NewZKCon(addrs []string) *ZKCon {
	r := new(ZKCon)
	r.lo = new(sync.Mutex)
	r.addrs = addrs

	return r
}

func (z *ZKCon) ConnectZK() error {
	con, _, err := zk.Connect(z.addrs, 20*time.Second)
	if err != nil {
		return err
	}

	z.con = con

	return nil
}

func (z *ZKCon) ReConnectZK() error {
	z.lo.Lock()
	defer z.lo.Unlock()

	z.CloseZK()

	return z.ConnectZK()
}

func (z *ZKCon) CloseZK() {
	if z.con != nil {
		z.con.Close()
	}
}

func (z *ZKCon) GetVal(path string) ([]byte, *zk.Stat, error) {
	z.lo.Lock()
	data, stat, err := z.con.Get(path)
	z.lo.Unlock()
	if err != nil {
		return nil, nil, err
	}
	return data, stat, nil
}

func (z *ZKCon) GetChilds(path string) ([]string, *zk.Stat, error) {
	z.lo.Lock()
	data, stat, err := z.con.Children(path)
	z.lo.Unlock()
	if err != nil {
		return nil, nil, err
	}

	return data, stat, nil
}

func (z *ZKCon) UpdateVal(pa string, val string) (*zk.Stat, error) {
	z.lo.Lock()
	stat, err := z.con.Set(pa, []byte(val), -1)
	z.lo.Unlock()
	if err != nil {
		return nil, err
	}

	return stat, nil
}

func (z *ZKCon) CreatePath(pa string, val string) (string, error) {
	z.lo.Lock()
	str,err := z.con.Create(pa, []byte(val), 0, zk.WorldACL(zk.PermAll))
	z.lo.Unlock()

	return str, err
}

func (z *ZKCon) DeletePath(pa string) error {
	z.lo.Lock()
	err := z.con.Delete(pa, -1)
	z.lo.Unlock()

	return err
}
