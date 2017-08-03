package dcache

import "errors"

var (
	cliPool     = GetCliPoolInst()
	sgm         = GetSGMInst()
	sgh         = GetSGHInst()
	msgQ        = GetMsgQInst()
	cleanUpFlag = GetCleanUpFlagInst()
)

var KeyNotExistError = errors.New("Key not exist.")

func Sync() {
	msgQ.Push("srvgroup", map[string]interface{}{
		"type": "sync",
	})
}
