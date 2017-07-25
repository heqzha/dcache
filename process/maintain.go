package process

import (
	"fmt"
	"time"

	"github.com/heqzha/dcache/core"
	"github.com/heqzha/dcache/utils"
	"github.com/heqzha/goutils/flow"
	"github.com/heqzha/goutils/logger"
)

var (
	fh      = flow.FlowNewHandler()
	conf    = utils.GetConfInst()
	msgQ    = utils.GetMsgQInst()
	sgm     = utils.GetSGMInst()
	sgh     = utils.GetSGHInst()
	cliPool = utils.GetCliPoolInst()
)

func MaintainSvrGroups() error {
	l, err := fh.NewLine(Receive, Handle, Reload)
	if err != nil {
		return err
	}
	fh.Start(l, flow.Params{})
	return nil
}

func StopAll() error {
	return fh.Destory()
}

func Receive(c *flow.Context) {
	for {
		msg := msgQ.Pop("srvgroup")
		if msg == nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		c.Set("msg", msg)
		c.Pass()
	}
}

func Handle(c *flow.Context) {
	msg := c.MustGet("msg").(map[string]interface{})
	c.Set("type", msg["type"])
	switch msg["type"].(string) {
	case "sync":
		fmt.Println("Handle: sync")
		dumps, err := sgm.Dump()
		if err != nil {
			logger.Error("SGM.Dump", err.Error())
			return
		}
		for gName, tb := range sgm.GetGroup() {
			for addr := range *tb {
				if conf.LocalAddr == addr {
					//Skip current service addr
					continue
				}
				c, err := cliPool.Get(addr)
				if err != nil {
					logger.Error(fmt.Sprintf("CSClientPool.Get: %s", addr), err.Error())
					return
				}
				res, err := c.SyncSrvGroup(dumps)
				if err != nil {
					logger.Error(fmt.Sprintf("CacheServClient.SyncSrvGroup: %s-%s", gName, addr), err.Error())
					return
				}
				if !res.GetStatus() {
					logger.Warn(fmt.Sprintf("CacheServClient.SyncSrvGroup: %s-%s", gName, addr), "Response status is false")
					return
				}
				tmpSGM := core.SGM{}
				tmpSGM.Load(res.GetSrvGroup())
				logger.Info("SGM.Merge", fmt.Sprintf("Before Merge: %s", sgm.CompareReadable(tmpSGM)))
				sgm.Merge(tmpSGM)
				logger.Info("SGM.Merge", fmt.Sprintf("After Merge: %s", sgm.CompareReadable(tmpSGM)))
				sgm.UpdateTimestamp(gName, addr)
			}
		}
		c.Next()
	case "ping":
		fmt.Println("Handle: ping")
	}
}

func Reload(c *flow.Context) {
	logger.Info("SGH.Load", "Reload group")
	sgh.Load(sgm.GetGroup())
}
