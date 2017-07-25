package test

import (
	"testing"
	"time"

	"github.com/heqzha/dcache/process"
	"github.com/heqzha/dcache/utils"
)

func TestMaintain(t *testing.T) {
	q := utils.GetMsgQInst()
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
