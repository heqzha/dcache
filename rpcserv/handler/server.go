package handler

import (
	"fmt"

	"github.com/heqzha/dcache/core"
	"github.com/heqzha/dcache/utils"
	"github.com/heqzha/goutils/logger"
)

var (
	sgm  = utils.GetSGMInst()
	msgQ = utils.GetMsgQInst()
)

func Set(group, key string, value []byte) error {
	//TODO
	return nil
}

func Get(group, key string) ([]byte, error) {
	//TODO
	return nil, nil
}

func Del(group, key string) ([]byte, error) {
	//TODO
	return nil, nil
}

func Register(group, addr string) error {
	logger.Debug("SGM.Register", fmt.Sprintf("%s: %s", group, addr))
	if err := sgm.Register(group, addr); err != nil {
		return err
	}
	msgQ.Push("srvgroup", map[string]interface{}{
		"type": "sync",
	})
	return nil
}

func Unregister(group, addr string) error {
	logger.Debug("SGM.Unregister", fmt.Sprintf("%s: %s", group, addr))
	if err := sgm.Unregister(group, addr); err != nil {
		return err
	}
	msgQ.Push("srvgroup", map[string]interface{}{
		"type": "sync",
	})
	return nil
}

func SyncSrvGroups(srvgroups []byte) (core.Condition, []byte, error) {
	tmpSGM := core.SGM{}
	tmpSGM.Load(srvgroups)
	logger.Debug("SGM.Merge", fmt.Sprintf("Before Merge: %s", sgm.CompareReadable(tmpSGM)))
	cond, err := sgm.Merge(tmpSGM)
	if err != nil {
		return -1, nil, err
	}
	logger.Debug("SGM.Merge", fmt.Sprintf("After Merge: %s", sgm.CompareReadable(tmpSGM)))
	dump, err := sgm.Dump()
	if err != nil {
		return -1, dump, err
	}
	return cond, dump, nil
}

func Ping(group, addr string) ([]byte, error) {
	logger.Debug("Ping", fmt.Sprintf("%s-%s", group, addr))
	dump, err := sgm.Dump()
	if err != nil {
		return nil, err
	}
	return dump, nil
}
