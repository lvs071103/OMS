package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oms/logic"
	"github.com/oms/models"
	"go.uber.org/zap"
)

// JenkinsInstancesListHandler Jenkins实例列表接口
// @Summary Jenkins实例列表接口
// @Description Jenkins实例列表接口
// @Tags Jenkins实例列表接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int false "页码"
// @Param page_size query int false "每页数据条数"
// @Success 200 {object} models.RespJenkinsInstancesList
// @Router /api/v1/app/release/jenkins/list [get]
func JenkinsInstancesListHandler(ctx *gin.Context) {
	// 解析参数
	page, pageSize := PaginationHander(ctx)
	// 获取列表
	data, err := logic.JenkinsInstancesList(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.JSONP(http.StatusOK, gin.H{"data": data, "message": "success"})
}

// JenkinsInstanceAddHandler 添加Jenkins实例
// @Summary 添加Jenkins实例
// @Description 添加Jenkins实例
// @Tags 添加Jenkins实例
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param release body models.CreateJenkinsInstanceRequest true "添加Jenkins实例"
// @Success 200 {object} map[string]interface{} "{"status": "ok"}"
// @Failure 500 {object} map[string]interface{} "{"error": "internal server error"}"
// @Failure 400 {object} map[string]interface{} "{"error": "invalid param"}"
// @Router /api/v1/app/release/jenkins/add [post]
func JenkinsInstanceAddHandler(ctx *gin.Context) {
	req := new(models.CreateJenkinsInstanceRequest)
	// 解析参数
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("JenkinsInstanceAdd with invalid param", zap.Error(err))
		ctx.JSON(StatusInvalidParam, gin.H{"error": err.Error()})
		return
	}

	// 添加Jenkins实例
	err := logic.JenkinsInstanceAdd(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
