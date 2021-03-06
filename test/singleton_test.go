package test

import (
	"fmt"
	"testing"

	"github.com/heqzha/dcache"
)

func TestMessageQueue(t *testing.T) {
	q := dcache.GetMsgQInst()
	for index := 0; index < 10; index++ {
		q.Push("test1", index)
	}

	for q.Len("test1") != 0 {
		fmt.Println(q.Pop("test1"))
	}
}
