package test

import (
	"testing"
	"time"

	"github.com/heqzha/dcache"
)

func TestDCacheString(t *testing.T) {
	pool := dcache.GetCliPoolInst()
	cli, err := pool.GetOrAdd("127.0.0.1:11000")
	if err != nil {
		t.Error(err)
		return
	}
	key := "test12"
	strVal := ""
	if err := cli.Get("default", key, &strVal); err != nil {
		t.Error(err)
		return
	}
	t.Log(strVal)
	if err := cli.Set("default", key, "Hello World"); err != nil {
		t.Error(err)
		return
	}

	if err := cli.Get("default", key, &strVal); err != nil {
		t.Error(err)
		return
	}
	t.Log(strVal)

	if err := cli.Del("default", key, &strVal); err != nil {
		t.Error(err)
		return
	}
	t.Log(strVal)
}

type TestObj struct {
	Name string
	Age  int
	Ts   time.Time
}

func TestDCacheObj(t *testing.T) {
	pool := dcache.GetCliPoolInst()
	cli, err := pool.GetOrAdd("127.0.0.1:11000")
	if err != nil {
		t.Error(err)
		return
	}
	key := "test1"

	obj := TestObj{
		Name: "abc",
		Age:  11,
		Ts:   time.Now(),
	}
	if err := cli.Set("default", key, obj); err != nil {
		t.Error(err)
		return
	}

	newObj := TestObj{}
	if err := cli.Get("default", key, &newObj); err != nil {
		t.Error(err)
		return
	}
	t.Log(newObj)

	delObj := TestObj{}
	if err := cli.Del("default", key, &delObj); err != nil {
		t.Error(err)
		return
	}
	t.Log(delObj)
}
