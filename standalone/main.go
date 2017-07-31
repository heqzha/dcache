package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
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
		//Notice: type of key is string and type of value is []byte, using utils.DCacheDecode to convert it to real type
		fmt.Println("Add:", key, value, reflect.TypeOf(value))
	}).LoaderFunc(func(key interface{}) (interface{}, error) {
		//Notice: type of key is string, the return value interface{} can be any type which will be convert to []byte for storage
		fmt.Println("Load:", key)
		return "Test", nil
	}).EvictedFunc(func(key, value interface{}) {
		//Notice: type of key is string and type of value is []byte, using utils.DCacheDecode to convert it to real type
		fmt.Println("Evicted:", key, value, reflect.TypeOf(value))
	}).Build()
	fmt.Printf("Start to Serving :%d with pid %d\n", conf.ServPort, pid)
	dcache.Run(conf.ServPort, cache)
}
