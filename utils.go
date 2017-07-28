package dcache

import (
	"github.com/heqzha/dcache/utils"
)

var (
	cliPool     = utils.GetCliPoolInst()
	sgm         = utils.GetSGMInst()
	sgh         = utils.GetSGHInst()
	msgQ        = utils.GetMsgQInst()
	cleanUpFlag = utils.GetCleanUpFlagInst()
)

func Sync() {
	msgQ.Push("srvgroup", map[string]interface{}{
		"type": "sync",
	})
}
