package test

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/heqzha/dcache"
)

var q = dcache.GetMsgQInst()

func TestCleanCloseInst(t *testing.T) {
	flg2 := dcache.GetCleanUpFlagInst()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	finish := false
	go func() {
		for _ = range signalChan {
			q.Push("test", "clean")
			finish = <-*flg2
			fmt.Println("closed")
		}
	}()

	go func() {
		for {
			s := q.Pop("test")
			if s == nil {
				time.Sleep(500 * time.Millisecond)
				continue
			}
			fmt.Println(s)
			clean()
		}
	}()
	fmt.Println("running")
	for !finish {
		time.Sleep(1 * time.Second)
	}
}

func clean() {
	flg1 := dcache.GetCleanUpFlagInst()
	*flg1 <- true
}
