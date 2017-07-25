package utils

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	ServPort   int    `yaml:"serv_port"`
	Debug      bool   `yaml:"debug"`
	LogDir     string `yaml:"log_dir"`
	IsRoot     bool   `yaml:"is_root"`
	RootAddr   string `yaml:"root_addr"`
	LocalGroup string `yaml:"local_group"`
	LocalAddr  string `yaml:"local_addr"`
}

func (c *Config) Init() {
	maxProcNum := flag.Int("maxproc", 0, "maximum number of CPUs")

	flag.Parse()
	if *maxProcNum == 0 {
		*maxProcNum = runtime.NumCPU()
	}
	runtime.GOMAXPROCS(*maxProcNum)

	p, _ := os.Getwd()
	cfile := filepath.Join(p, "./config.yml")

	absCFile, err := filepath.Abs(cfile)
	if err != nil {
		log.Printf("No correct config file: %s - %s", cfile, err.Error())
		os.Exit(1)
	}
	buf, err := ioutil.ReadFile(absCFile)
	if err != nil {
		log.Printf("Failed to read config file <%s> : %s", absCFile, err.Error())
		os.Exit(1)
	}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		log.Printf("Failed to parse config fliel <%s> : %s", absCFile, err.Error())
		os.Exit(1)
	}
}
