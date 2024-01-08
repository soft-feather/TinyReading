package meta

import "github.com/gin-gonic/gin"

func SetViewRouterGroup(router *gin.RouterGroup) {
	searchRouteGroup := router.Group("/meta")
	v1SearchRouteGroup := searchRouteGroup.Group("/v1")
	{
		v1SearchRouteGroup.GET("/", SearchHandler)
	}
}
