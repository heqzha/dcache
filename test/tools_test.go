package test

import (
	"testing"
	"time"

	"github.com/heqzha/dcache/utils"
)

type TestStruct struct {
	Name string
	Age  int
	Ts   time.Time
}

func TestDCacheEncode(t *testing.T) {

	data, err := utils.DCacheEncode("string")
	if err != nil {
		t.Error(err)
		return
	}
	strVal := ""
	if err := utils.DCacheDecode(&strVal, data); err != nil {
		t.Error(err)
		return
	}
	t.Log(strVal)

	obj := TestStruct{
		Name: "abc",
		Age:  12,
		Ts:   time.Now(),
	}
	data, err = utils.DCacheEncode(obj)
	if err != nil {
		t.Error(err)
		return
	}
	newObj := TestStruct{}
	if err := utils.DCacheDecode(&newObj, data); err != nil {
		t.Error(err)
		return
	}
	t.Log(newObj)
}
