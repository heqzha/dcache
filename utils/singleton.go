package utils

import (
	"sync"

	"github.com/heqzha/dcache/core"
	"github.com/heqzha/dcache/rpccli"
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
var cliPoolInst rpccli.CSClientPool
var cliPoolInstOnce sync.Once

func GetCliPoolInst() rpccli.CSClientPool {
	cliPoolInstOnce.Do(func() {
		cliPoolInst = make(rpccli.CSClientPool)
	})
	return cliPoolInst
}
