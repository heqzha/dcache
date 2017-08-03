package test

import (
	"fmt"
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

func TestDCacheObjGetSet(t *testing.T) {
	pool := dcache.GetCliPoolInst()
	cli, err := pool.GetOrAdd("127.0.0.1:11000")
	if err != nil {
		t.Error(err)
		return
	}
	key := "test1"

	oldObj := TestObj{}
	if err := cli.Get("default", key, &oldObj); err != nil {
		t.Error(err)
		return
	}
	t.Log(oldObj)

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
}

func TestDCacheObjDel(t *testing.T) {
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

	delObj := TestObj{}
	if err := cli.Del("default", key, &delObj); err != nil {
		t.Error(err)
		return
	}
	t.Log(delObj)
}

func TestDCacheObjGetIfExist(t *testing.T) {
	pool := dcache.GetCliPoolInst()
	cli, err := pool.GetOrAdd("127.0.0.1:11000")
	if err != nil {
		t.Error(err)
		return
	}
	group := "default"
	key := "key1"
	obj := TestObj{
		Name: "abc",
		Age:  11,
		Ts:   time.Now(),
	}
	if err := cli.SetWithExpire(group, key, obj, 10*time.Second); err != nil {
		t.Error(err)
		return
	}
	for {
		newObj := TestObj{}
		if err := cli.GetIfExist(group, key, &newObj); err != nil {
			if err == dcache.KeyNotExistError {
				fmt.Println("Done")
				break
			} else {
				t.Error(err)
				return
			}
		}
		fmt.Println(newObj)
		time.Sleep(time.Second)
	}
}
