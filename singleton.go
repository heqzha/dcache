package dcache

import (
	"sync"

	"github.com/heqzha/dcache/core"
)

// SGM instance
var sgmInst *core.SGM
var sgmInstOnce sync.Once

func GetSGMInst() *core.SGM {
	sgmInstOnce.Do(func() {
		sgmInst = new(core.SGM)
		sgmInst.Init()
	})
	return sgmInst
}

// SGHash instance
var sghInst *core.SGHash
var sghInstOnce sync.Once

func GetSGHInst() *core.SGHash {
	sghInstOnce.Do(func() {
		sghInst = new(core.SGHash)
		sghInst.Init(3, nil)
	})
	return sghInst
}

//Message queue instance
var msgQInst *core.MessageQueue
var msgQInstOnce sync.Once

func GetMsgQInst() *core.MessageQueue {
	msgQInstOnce.Do(func() {
		msgQInst = new(core.MessageQueue)
		msgQInst.Init()
	})
	return msgQInst
}

//RPC client pool instance
var cliPoolInst *CSClientPool
var cliPoolInstOnce sync.Once

func GetCliPoolInst() *CSClientPool {
	cliPoolInstOnce.Do(func() {
		cliPoolInst = new(CSClientPool)
		cliPoolInst.Init()
	})
	return cliPoolInst
}

var cleanUpFlagInst *chan bool
var cleanUpFlagInstOnce sync.Once

func GetCleanUpFlagInst() *chan bool {
	cleanUpFlagInstOnce.Do(func() {
		inst := make(chan bool)
		cleanUpFlagInst = &inst
	})
	return cleanUpFlagInst
}
