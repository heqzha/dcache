package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/heqzha/dcache"
	"github.com/heqzha/goutils/logger"
)

var (
	conf    *Config
	cliPool = dcache.GetCliPoolInst()
)

func init() {
	conf = new(Config)
	conf.Init()
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
	cache := dcache.New(1024).Simple().IsRoot(conf.IsRoot).RootAddr(conf.RootAddr).LocalAddr(conf.LocalAddr).LocalGroup(conf.LocalGroup).AddedFunc(func(key, value interface{}) {
		fmt.Printf("Add: %v-%v\n", key, string(value.([]byte)))
	}).LoaderFunc(func(key interface{}) (interface{}, error) {
		//TODO
		fmt.Printf("Load: %v\n", key)
		return []byte("Test"), nil
	}).EvictedFunc(func(key, value interface{}) {
		//TODO
		fmt.Printf("Evicted: %v-%v\n", key, string(value.([]byte)))
	}).Build()
	fmt.Printf("Start to Serving :%d with pid %d\n", conf.ServPort, pid)
	dcache.Run(conf.ServPort, cache)
}
