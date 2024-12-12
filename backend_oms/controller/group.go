package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oms/logic"
	"github.com/oms/models"
	"go.uber.org/zap"
)

// GroupListHander 用户列表接口
// @Summary 用户组列表接口
// @Description 展示用户组列表接口
// @Tags 用户组列表接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int false "页码"
// @Param page_size query int false "每页数据条数"
// @Success 200 {object} models.RespAuthGroup
// @Router /api/v1/group/list [get]
func GroupListHander(ctx *gin.Context) {
	page, pageSize := PaginationHander(ctx)
	data, err := logic.GroupList(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

// GroupAddHandler - 添加用户组
func GroupAddHandler(ctx *gin.Context) {
	group := new(models.CreateGroupRequest)
	// ShouldBindJSON 解析请求参数
	if err := ctx.ShouldBindJSON(group); err != nil {
		zap.L().Error("GroupAdd with invalid param", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := logic.AddGroupLogic(group)
	if err != nil {
		zap.L().Error("logic.AddGroupLogic failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ResponseSuccess(ctx, nil)
	ctx.JSON(http.StatusOK, gin.H{"message": "添加成功"})
}

// GroupDetailHandler 组详情
func GroupDetailHandler(ctx *gin.Context) {
	groupID := ctx.Param("id")
	id, err := strconv.ParseInt(groupID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := logic.GroupDetail(id)
	if err != nil {
		zap.L().Error("logic.GroupDetail failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

// GroupUpdateHandler 更新组
func GroupUpdateHandler(ctx *gin.Context) {
	groupID := ctx.Param("id")
	id, err := strconv.ParseInt(groupID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req := new(models.CreateGroupRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("GroupUpdate with invalid param", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = logic.GroupUpdate(id, req)
	if err != nil {
		zap.L().Error("logic.GroupUpdate failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 编辑成功
	ctx.JSON(http.StatusOK, gin.H{"message": "更新成功"})

}

// GroupDeleteHandler 删除组
func GroupDeleteHandler(ctx *gin.Context) {
	groupID := ctx.Param("id")
	id, err := strconv.ParseInt(groupID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = logic.GroupDelete(id)
	if err != nil {
		zap.L().Error("logic.GroupDelete failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
