package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oms/logic"
	"github.com/oms/models"
)

// PermissionListHander - 权限列表
func PermissionListHander(ctx *gin.Context) {
	// 获取列表
	data, err := logic.PermissionList()
	if err != nil {
		// ResponseError(ctx, CodeServerBusy)
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// ResponseSuccess(ctx, data)
	ctx.JSON(http.StatusOK, gin.H{"data": data, "message": "success"})
}

// PermissionAddHandler - 添加权限
func PermissionAddHandler(ctx *gin.Context) {
	// 1. 参数校验
	permission := new(models.CreatePermissionRequest)
	if err := ctx.ShouldBindJSON(permission); err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	// 2. 业务处理
	err := logic.PermissionAdd(permission)
	if err != nil {
		ResponseError(ctx, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(ctx, nil)
}
