package dcache

import "github.com/heqzha/dcache/global"

var (
	cliPool     = global.GetCliPoolInst()
	sgm         = global.GetSGMInst()
	sgh         = global.GetSGHInst()
	msgQ        = global.GetMsgQInst()
	cleanUpFlag = global.GetCleanUpFlagInst()
)

func Sync() {
	msgQ.Push("srvgroup", map[string]interface{}{
		"type": "sync",
	})
}
