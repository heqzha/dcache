package test

import (
	"testing"

	"github.com/heqzha/dcache/utils"
)

func TestDCache(t *testing.T) {
	pool := utils.GetCliPoolInst()
	cli, err := pool.Get("127.0.0.1:11000")
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
