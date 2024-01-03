package view

import "github.com/gin-gonic/gin"

func SetViewRouterGroup(router *gin.RouterGroup) {
	viewRouteGroup := router.Group("/view")
	v1viewRouteGroup := viewRouteGroup.Group("/v1")

	{
		v1viewRouteGroup.GET("/book/:id", BookViewHandler)
	}
}
