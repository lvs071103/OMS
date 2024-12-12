package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "UserID"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUserID 获取当前登录的用户ID
func getCurrentUserID(ctx *gin.Context) (userID int64, err error) {
	uid, ok := ctx.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func PaginationHander(ctx *gin.Context) (page, pageSize int64) {
	// 获取分页参数
	pageStr := ctx.Query("page")
	pageSizeStr := ctx.Query("pageSize")
	page, pageErr := strconv.ParseInt(pageStr, 10, 64)
	if pageErr != nil {
		page = 1
	}
	pageSize, pageSizeErr := strconv.ParseInt(pageSizeStr, 10, 64)
	if pageSizeErr != nil {
		pageSize = 10
	}
	return
}
