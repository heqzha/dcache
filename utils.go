package dcache

var (
	cliPool     = GetCliPoolInst()
	sgm         = GetSGMInst()
	sgh         = GetSGHInst()
	msgQ        = GetMsgQInst()
	cleanUpFlag = GetCleanUpFlagInst()
)

func Sync() {
	msgQ.Push("srvgroup", map[string]interface{}{
		"type": "sync",
	})
}
