package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oms/logic"
	"github.com/oms/models"
	"go.uber.org/zap"
)

// EnvListHandler - 环境列表
func EnvListHandler(ctx *gin.Context) {
	page, pageSize := PaginationHander(ctx)
	data, err := logic.EnvList(page, pageSize)
	if err != nil {
		// ResponseError(ctx, CodeServerBusy)
		zap.L().Error("logic.EnvList failed", zap.Error(err))
		return
	}
	// ResponseSuccess(ctx, data)
	ctx.JSON(http.StatusOK, data)
}

// EnvAddHandler - 添加环境
func EnvAddHandler(ctx *gin.Context) {
	env := new(models.CreateEnvRequest)
	if err := ctx.ShouldBindJSON(env); err != nil {
		zap.L().Error("EnvAdd with invalid param", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	err := logic.EnvAddHandler(env)
	if err != nil {
		zap.L().Error("logic.EnvAddHandler failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, nil)
}

// EnvDetailHandler - 环境详情
func EnvDetailHandler(ctx *gin.Context) {
	envID := ctx.Param("id")
	id, err := strconv.ParseInt(envID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := logic.EnvDetail(id)
	if err != nil {
		zap.L().Error("logic.EnvDetail failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

// EnvUpdateHandler - 更新环境
func EnvUpdateHandler(ctx *gin.Context) {
	envID := ctx.Param("id")
	id, err := strconv.ParseInt(envID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 参数校验
	req := new(models.CreateEnvRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("EnvUpdate with invalid param", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 业务处理
	err = logic.EnvUpdate(id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// EnvDeleteHandler - 删除环境
func EnvDeleteHandler(ctx *gin.Context) {
	// 参数校验
	envID := ctx.Param("id")
	id, err := strconv.ParseInt(envID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 业务处理
	err = logic.EnvDelete(id)
	if err != nil {
		zap.L().Error("logic.EnvDelete failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
