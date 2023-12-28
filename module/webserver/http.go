package webserver

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"os"
	"syscall"
	"time"
)

const (
	DefaultTimeoutSecond = 5
)

var (
	DefaultWebserver             *Webserver
	DefaultWebserverShutdownChan chan os.Signal
)

//const (
//	DefaultAddress = "0.0.0.0:5000"
//)

type Webserver struct {
	server       *http.Server
	shutdownChan chan os.Signal
}

func (w *Webserver) Init(address string) error {
	logger.Info(address)

	w.shutdownChan = make(chan os.Signal)
	r := gin.Default()
	r.Use(cors.Default())
	// TODO logic route group addition
	//rg := r.Group("/")
	//rg.POST("/test", func(c *gin.Context) {
	//	logger.Info("test")
	//  ctx.Json({})
	//	return
	//})
	r.GET("/", func(c *gin.Context) {
		logger.Info("test /")
		c.JSON(200, "")
	})
	// api.SetRouterGroup(r)

	var err error

	go func() {
		w.server = &http.Server{Addr: address, Handler: r.Handler()}
		err = w.server.ListenAndServe()
	}()

	go func() {
		select {
		case signal := <-w.shutdownChan:
			switch signal {
			case syscall.SIGINT:
				logger.Info("webserver shutdown ...")
				ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeoutSecond*time.Second)
				defer cancel()
				if err = w.server.Shutdown(ctx); err != nil {
					logger.Fatal("shutdown %v", err)
				}
			}
		}
	}()

	if err != nil {
		logger.Fatal("webserver init: %v", err)
		return err
	}
	return err
}
