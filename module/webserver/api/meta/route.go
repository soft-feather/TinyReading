package meta

import "github.com/gin-gonic/gin"

func SetMetaRouterGroup(router *gin.RouterGroup) {
	metaRouteGroup := router.Group("/meta")
	v1MetaRouteGroup := metaRouteGroup.Group("/v1")
	{
		v1MetaRouteGroup.GET("/search", SearchHandler)
		v1MetaRouteGroup.POST("/add", AddHandler)
	}
}
