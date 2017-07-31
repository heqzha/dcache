package test

import (
	"testing"
	"time"

	"github.com/heqzha/dcache"
)

func TestMaintain(t *testing.T) {
	q := dcache.GetMsgQInst()
	for index := 0; index < 5; index++ {
		q.Push("srvgroup", map[string]interface{}{
			"type": "sync",
		})
	}
	for index := 0; index < 3; index++ {
		q.Push("srvgroup", map[string]interface{}{
			"type": "ping",
		})
	}

	dcache.ProcMaintainSvrGroups()
	time.Sleep(time.Second * 5)
}
