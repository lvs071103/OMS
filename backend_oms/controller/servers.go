package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oms/logic"
)

// ServerListHandler 服务器列表接口
// @Summary 服务器列表接口
// @Description 服务器列表接口
// @Tags 服务器列表接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "
// @Param page query int false "页码"
// @Param page_size query int false "每页数据条数"
// @Success 200 {object} models.RespServerList
// @Router /api/v1/app/release/server/list [get]
func ServerListHandler(ctx *gin.Context) {
	// 解析参数
	page, pageSize := PaginationHander(ctx)
	// 获取列表
	data, err := logic.ServerList(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.JSONP(http.StatusOK, gin.H{"data": data, "message": "success"})
}
