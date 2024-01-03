package api

import (
	"github.com/gin-gonic/gin"
	"github.com/soft-feather/TinyReading/module/webserver/api/view"
)

func SetRouterGroup(router *gin.Engine) {
	ApiRouteGroup := router.Group("/api")
	view.SetViewRouterGroup(ApiRouteGroup)
}
