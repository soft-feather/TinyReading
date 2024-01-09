package meta

import (
	"github.com/gin-gonic/gin"
	"github.com/soft-feather/TinyReading/util/errors"
	"github.com/soft-feather/TinyReading/util/http"
	"github.com/soft-feather/TinyReading/util/storage/sqlite"
	"github.com/soft-feather/TinyReading/util/storage/sqlite/table"
	"github.com/wonderivan/logger"
	"time"
)

type SearchMetaRequest struct {
	Filename         string
	SearchPreviousAt time.Time
	SearchAfterAt    time.Time
}

type AddMetaRequest struct {
	BookId          uint
	BookName        string
	BookAuthor      string
	BookDescription string
	BookSN          string

	BookTags             []table.Tag
	BookPublishDate      time.Time
	BookMetaCreateAt     time.Time
	BookMetaLastModifyAt time.Time
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

// AddHandler add book meta handler
// curl -X POST -H "Content-Type: application/json" -d '{"book_id": 1, "book_name": "test", "book_author": "test", "book_description": "test", "book_sn": "test", "book_tags": [{"name": "test", "nice": 1}], "book_publish_date": "2020-01-01T00:00:00Z", "book_meta_create_at": "2020-01-01T00:00:00Z", "book_meta_last_modify_at": "2020-01-01T00:00:00Z"}' 127.0.0.1:5000/api/meta/add
func AddHandler(ctx *gin.Context) {
	req := AddMetaRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		http.Responses(ctx, errors.ValidationException, nil)
		return
	}
	logger.Debug("req: ", req)
	data := &table.Meta{
		ID:           req.BookId,
		BookSN:       req.BookSN,
		PublishDate:  req.BookPublishDate,
		CreateAt:     req.BookMetaCreateAt,
		LastModifyAt: req.BookMetaLastModifyAt,
		Author:       req.BookAuthor,
		Name:         req.BookName,
		Description:  req.BookDescription,
	}
	// TODO tag ls
	db := sqlite.GetDB().Create(data)
	if db.Error != nil {
		logger.Error(db.Error)
		http.Responses(ctx, errors.HandleInternalException, nil)
		return
	}
	http.Responses(ctx, errors.OK, data)
	//http.Responses(ctx, errors.NotImplement, nil)
}
