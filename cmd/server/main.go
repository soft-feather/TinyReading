package main

import (
	"flag"
	"github.com/soft-feather/TinyReading/module/webserver"
	"github.com/soft-feather/TinyReading/util/config"
	"github.com/soft-feather/TinyReading/util/log"
	"os"
	"os/signal"
	"syscall"
)

const (
	DefaultShutdownSecond = 1
	DefaultDryRunSecond   = 2
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
	}

	//if dryRun {
	//	go func() {
	//		// 在这里执行 dry-run 逻辑
	//		// 例如，可以打印出将要执行的操作但不实际执行
	//		time.Sleep(time.Second * DefaultDryRunSecond)
	//		logger.Info("Dry-run mode enabled, no operations will be performed")
	//		webserver.GetDefaultWebserverShutdownChan() <- syscall.SIGINT
	//	}()
	//}
	//
	//if !dryRun {
	//	webserver.GetDefaultWebserverShutdownChan() <- syscall.SIGINT
	//	time.Sleep(time.Second * DefaultShutdownSecond)
	//}
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
