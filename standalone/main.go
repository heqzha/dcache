package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/heqzha/dcache/process"
	"github.com/heqzha/dcache/rpcserv"
	"github.com/heqzha/dcache/utils"

	"github.com/heqzha/goutils/logger"
)

var (
	conf    *utils.Config
	cliPool = utils.GetCliPoolInst()
)

func init() {
	conf = utils.GetConfInst()
	logger.Config(conf.LogDir, logger.LOG_LEVEL_DEBUG)
}

func CreatePID(name string) int {
	wd, _ := os.Getwd()
	pidFile, err := os.OpenFile(filepath.Join(wd, fmt.Sprintf("%s.pid", name)), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("failed to create pid file: %s", err.Error())
		os.Exit(1)
	}
	defer pidFile.Close()

	pid := os.Getpid()
	pidFile.WriteString(strconv.Itoa(pid))
	return pid
}

func main() {
	pid := CreatePID("dcache")
	process.MaintainSvrGroups()
	defer process.StopAll()

	if !conf.IsRoot {
		root, err := cliPool.Add(conf.RootAddr)
		if err != nil {
			panic(fmt.Sprintf("failed to connect root node: %s", err.Error()))
		}
		res, err := root.Register(conf.LocalGroup, conf.LocalAddr)
		if err != nil {
			panic(fmt.Sprintf("failed to register to root node: %s", err.Error()))
		} else if !res.GetStatus() {
			panic(fmt.Sprintf("register denied"))
		}
	}

	fmt.Printf("Start to Serving :%d with pid %d\n", conf.ServPort, pid)
	rpcserv.Run(conf.ServPort)
}
