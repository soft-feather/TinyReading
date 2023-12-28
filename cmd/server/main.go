package main

import (
	"github.com/soft-feather/TinyReading/module/webserver"
	"github.com/soft-feather/TinyReading/util/config"
	"github.com/soft-feather/TinyReading/util/log"
	"os"
)

func main() {
	var err error
	err = Init()
	if err != nil {
		os.Exit(1)
	}
	err = webserver.InitDefaultWebserver()
	if err != nil {
		os.Exit(1)
	}
	select {}
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
