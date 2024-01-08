package meta

import (
	"github.com/gin-gonic/gin"
	"github.com/soft-feather/TinyReading/util/errors"
	"github.com/soft-feather/TinyReading/util/http"
	"github.com/wonderivan/logger"
	"time"
)

type SearchMetaRequest struct {
	Filename         string
	SearchPreviousAt time.Time
	SearchAfterAt    time.Time
}

// SearchHandler 获取书籍视图，返回相应的书籍文件流
func SearchHandler(ctx *gin.Context) {
	// 根据 filename 等信息查询书籍 id
	// 获取书籍 绝对地址
	// 返回书籍文件
	req := SearchMetaRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		http.Responses(ctx, errors.ValidationException, nil)
		return
	}
	logger.Debug("req: ", req)
	if req.Filename == "" {
		http.Responses(ctx, errors.IdValidationException, "暂不支持搜索空文件名")
		return
	}
	// db search
	//testFile := path.Join(profile.GetEtcProfile().CWD, "doc/pdf", "test.pdf")
	//logger.Debug("test file: ", testFile)
	//ctx.File(testFile)
}
