package utils

import (
	"sync"

	"github.com/bluele/gcache"
	"github.com/heqzha/dcache/core"
	"github.com/heqzha/dcache/rpccli"
)

// Config instance
var confInst *Config
var confInstOnce sync.Once

func GetConfInst() *Config {
	confInstOnce.Do(func() {
		confInst = new(Config)
		confInst.Init()
	})
	return confInst
}

// SGM instance
var sgmInst *core.SGM
var sgmInstOnce sync.Once

func GetSGMInst() *core.SGM {
	conf := GetConfInst()
	sgmInstOnce.Do(func() {
		if conf.LocalAddr == "" {
			panic("missing addr in config.yml")
		}
		sgmInst = new(core.SGM)
		sgmInst.Init(conf.LocalGroup, conf.LocalAddr)
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

//Cache instance
var cacheInst *gcache.Cache
var cacheInstOnce sync.Once

func GetCacheInst() *gcache.Cache {
	cacheInstOnce.Do(func() {
		c := gcache.New(1024).LFU().AddedFunc(func(key, value interface{}) {
			//TODO
		}).LoaderFunc(func(key interface{}) (interface{}, error) {
			//TODO
			return nil, nil
		}).EvictedFunc(func(key, value interface{}) {
			//TODO
		}).Build()
		cacheInst = &c
	})
	return cacheInst
}
