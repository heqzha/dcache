package dcache

import (
	"fmt"
	"time"

	"github.com/heqzha/dcache/core"
	"github.com/heqzha/goutils/flow"
	"github.com/heqzha/goutils/logger"
)

var (
	fh = flow.FlowNewHandler()
)

func ProcMaintainSvrGroups() error {
	l, err := fh.NewLine(receive, handle, reload)
	if err != nil {
		return err
	}
	fh.Start(l, flow.Params{})
	return nil
}

func ProcStopAll() error {
	return fh.Destory()
}

func receive(c *flow.Context) {
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

func handle(c *flow.Context) {
	msg := c.MustGet("msg").(map[string]interface{})
	c.Set("type", msg["type"])
	switch msg["type"].(string) {
	case "sync":
		dumps, err := sgm.Dump()
		if err != nil {
			logger.Error("SGM.Dump", err.Error())
			return
		}
		for gName, tb := range sgm.GetGroup() {
			for addr := range *tb {
				if sgm.GetLocalAddr() == addr {
					//Skip local addr
					continue
				}
				c, err := cliPool.GetOrAdd(addr)
				if err != nil {
					logger.Error(fmt.Sprintf("CSClientPool.Get: %s", addr), err.Error())
					return
				}
				res, err := c.syncSrvGroup(dumps)
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
	case "unregister":
		for gName, tb := range sgm.GetGroup() {
			for addr := range *tb {
				if sgm.GetLocalAddr() == addr {
					//Skip local addr
					continue
				}
				c, err := cliPool.GetOrAdd(addr)
				if err != nil {
					logger.Error(fmt.Sprintf("CSClientPool.Get: %s", addr), err.Error())
					return
				}
				res, err := c.unregister(sgm.GetLocalGroup(), sgm.GetLocalAddr())
				if err != nil {
					logger.Error(fmt.Sprintf("CacheServClient.Unregister: %s-%s", gName, addr), err.Error())
					return
				}
				if !res.GetStatus() {
					logger.Warn(fmt.Sprintf("CacheServClient.Unregister: %s-%s", gName, addr), "Response status is false")
					return
				}
			}
		}
		*cleanUpFlag <- true
	}
}

func reload(c *flow.Context) {
	logger.Info("SGH.Load", "Reload group")
	sgh.Load(sgm.GetGroup())
}
