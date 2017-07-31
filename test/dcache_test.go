package test

import (
	"testing"

	"github.com/heqzha/dcache/global"
)

func TestDCache(t *testing.T) {
	pool := global.GetCliPoolInst()
	cli, err := pool.GetOrAdd("127.0.0.1:11000")
	if err != nil {
		t.Error(err)
		return
	}
	key := "test12"
	getRes, err := cli.Get("default", key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(getRes.GetValue()))
	if _, err := cli.Set("default", key, []byte("Hello World")); err != nil {
		t.Error(err)
		return
	}
	getRes, err = cli.Get("default", key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(getRes.GetValue()))

	delRes, err := cli.Del("default", key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(delRes.GetValue()))
}
