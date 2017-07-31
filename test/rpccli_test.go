package test

import (
	"testing"

	"github.com/heqzha/dcache/core"
	"github.com/heqzha/dcache/global"
)

func TestRPCClient(t *testing.T) {
	pool := global.GetCliPoolInst()
	cli, err := pool.GetOrAdd("127.0.0.1:11000")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := cli.Register("test1", "127.0.0.1:11001")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("cli.Register: %t", res.Status)
}

func TestRPCClientPing(t *testing.T) {
	pool := global.GetCliPoolInst()
	cli, err := pool.GetOrAdd("127.0.0.1:11000")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := cli.Ping("test", "127.0.0.1:11009")
	if err != nil {
		t.Error(err)
		return
	}
	sgm := new(core.SGM)
	sgm.Init()
	sgm.RegisterLocalAddr("test", "127.0.0.1:11009")
	sgm.Load(res.SrvGroup)
	table, err := sgm.GetTable("default")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(table.String())
}
