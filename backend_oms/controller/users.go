package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/oms/dao/mysql"
	"github.com/oms/logic"
	"github.com/oms/models"
	"go.uber.org/zap"
)

// UserListHander 用户管理列
func UserListHander(ctx *gin.Context) {
	page, pageSize := PaginationHander(ctx)
	data, err := logic.UserList(page, pageSize)
	if err != nil {
		zap.L().Error("logic.UserList failed.", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UserDetailHandler 用户详情页
func UserDetailHandler(ctx *gin.Context) {
	// 1.获取参数从URL中获取贴子的id
	userID := ctx.Param("id")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		// ResponseError(ctx, CodeInvalidParam)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 2.根据id取出用户的详细信息
	data, err := logic.UserDetail(id)
	if err != nil {
		zap.L().Error("logic.PostDetail() failed", zap.Error(err))
		// ResponseError(ctx, CodeServerBusy)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 3.响应数据
	// ResponseSuccess(ctx, data)
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

// UserUpdateHandler 用户更新
func UserUpdateHandler(ctx *gin.Context) {
	// 1. 参数校验
	userID := ctx.Param("id")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}

	userParams := new(models.CreateUserRequest)
	if err := ctx.ShouldBindJSON(userParams); err != nil {
		// 请求参数有误, 直接返回响应
		zap.L().Error("UserUpdate with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 更新用户
	if err := logic.UserUpdate(id, userParams); err != nil {
		zap.L().Error("logic.UserUpdate failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			// ResponseError(ctx, CodeUserExists)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	// 3. 返回响应
	ResponseSuccess(ctx, nil)
}

// UserAddHandler 后台添加用户
func UserAddHandler(ctx *gin.Context) {
	// 1. 参数校验
	userParams := new(models.CreateUserRequest)
	if err := ctx.ShouldBindJSON(userParams); err != nil {
		// 请求参数有误, 直接返回响应
		zap.L().Error("UserUpdate with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// ResponseError(ctx, CodeInvalidParam)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": removeTopStruct(errs.Translate(trans))})
		return
	}

	fmt.Println(userParams)

	// 2. 处理添加逻辑
	if err := logic.UserAdd(userParams); err != nil {
		zap.L().Error("logic.UserAdd failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			// ResponseError(ctx, CodeUserExists)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	// 返回
	// ResponseSuccess(ctx, nil)
	ctx.JSON(http.StatusOK, gin.H{"message": "添加成功"})
}

// UserDeleteHandler - 删除用户
func UserDeleteHandler(ctx *gin.Context) {
	// 1. 参数校验
	userID := ctx.Param("id")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		// ResponseError(ctx, CodeInvalidParam)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := logic.UserDelete(id); err != nil {
		zap.L().Error("logic.UserDelete failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(ctx, CodeUserNotExists)
		}
		return
	}
	ResponseSuccess(ctx, nil)
}
