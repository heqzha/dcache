package test

import (
	"testing"
	"time"

	"github.com/heqzha/dcache/global"
	"github.com/heqzha/dcache/process"
)

func TestMaintain(t *testing.T) {
	q := global.GetMsgQInst()
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

	process.MaintainSvrGroups()
	time.Sleep(time.Second * 5)
}
