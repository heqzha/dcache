package dcache

import "github.com/heqzha/dcache/utils"

var (
	cliPool = utils.GetCliPoolInst()
	sgm     = utils.GetSGMInst()
	sgh     = utils.GetSGHInst()
	msgQ    = utils.GetMsgQInst()
)

func Sync() {
	msgQ.Push("srvgroup", map[string]interface{}{
		"type": "sync",
	})
}
