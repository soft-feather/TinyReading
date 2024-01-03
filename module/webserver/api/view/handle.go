package view

import "github.com/gin-gonic/gin"

// BookViewHandler 获取书籍视图，返回相应的书籍文件流
func BookViewHandler(ctx *gin.Context) {

	// TODO
	// 根据 id 查询
	// 获取书籍 绝对地址
	// 返回书籍文件
	ctx.File("")
}
