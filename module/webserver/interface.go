package webserver

import (
	"github.com/soft-feather/TinyReading/util/config"
	"os"
)

// 使用说明
// InitDefaultWebserver 在只需要实现全局单例时使用
// 全局单例下调用	webserver.GetDefaultWebserverShutdownChan() <- syscall.SIGINT 则 webserver shutdown

// GetDefaultWebserverShutdownChan
// shutdown webserver
func GetDefaultWebserverShutdownChan() chan os.Signal {
	return DefaultWebserverShutdownChan
}

// InitDefaultWebserver
// 初始化默认的 webserver
func InitDefaultWebserver() error {
	DefaultWebserver = &Webserver{}
	addrLs := config.GetConfig().Default.AddrLs
	err := DefaultWebserver.Init(addrLs)
	DefaultWebserverShutdownChan = DefaultWebserver.shutdownChan
	return err
}
