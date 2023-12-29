package main

import (
	"flag"
	"github.com/soft-feather/TinyReading/module/webserver"
	"github.com/soft-feather/TinyReading/util/config"
	"github.com/soft-feather/TinyReading/util/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	DefaultShutdownSecond = 1
)

func main() {

	var dryRun bool

	flag.BoolVar(&dryRun, "dry-run", false, "Enable dry-run mode")
	flag.Parse()
	var err error
	err = Init()
	if err != nil {
		os.Exit(1)
	}
	err = webserver.InitDefaultWebserver()
	if err != nil {
		os.Exit(1)
	}

	if dryRun {
		webserver.GetDefaultWebserverShutdownChan() <- syscall.SIGINT
	} else {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
		// 等待信号
		<-signalChan
		time.Sleep(time.Second * DefaultShutdownSecond)
	}
}

func Init() error {
	var err error
	err = log.Init()
	if err != nil {
		return err
	}
	err = config.Init()
	if err != nil {
		return err
	}
	return nil
}
