package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/oms/dao/mysql"
	"github.com/oms/logic"
	"github.com/oms/models"
	"go.uber.org/zap"
)

// 处理注册请求的函数
func SignUpHandler(ctx *gin.Context) {
	// 1. 参数校验
	p := new(models.ParamSignUp)
	if err := ctx.ShouldBindJSON(p); err != nil {
		// 请求参数有误, 直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		// errs, ok := err.(validator.ValidationErrors)
		// if !ok {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		// 翻译错误
		// ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 2. 验证验证码
	if !logic.VerifyCode(p.Email, p.VerifyCode) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "验证码错误"})
		return
	}

	if p.Password != p.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "两次输入的密码不一致"})
		return
	}

	// 3. 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		// if errors.Is(err, mysql.ErrorUserExist) {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// }
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 4. 返回响应
	ctx.JSON(http.StatusOK, nil)
}

func LoginHandler(ctx *gin.Context) {
	// 1.获取请求参数
	p := new(models.ParamsLogin)
	// 参数校验
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": removeTopStruct(errs.Translate(trans))})
		return
	}
	// 2.业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.UserName), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// ResponseError(ctx, CodeInvalidPassword)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 3.返回响应
	ResponseSuccess(ctx, gin.H{
		"uid": user.ID,
		// 前端js: id值大于1<<53-1;后端：int64类型的最大值1<<64-1，超出有失真问题
		"username": user.UserName,
		"token":    user.Token,
	})
}

// LogoutHandler 处理用户登出请求
func LogoutHandler(ctx *gin.Context) {
	// 清除客户端的身份验证信息
	// 如果使用的是 JWT，可以让客户端删除存储的令牌
	// 如果使用的是会话，可以清除会话信息

	// 这里假设使用的是 JWT，客户端需要删除存储的令牌
	// 你可以在这里添加任何需要的逻辑，例如记录登出日志等

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}

// GenerateCaptchaHandler - 生成验证码
func GenerateCaptchaHandler(ctx *gin.Context) {
	// 1. 获取参数
	req := new(models.CaptchaRequest)
	// 2. 参数校验
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("GenerateCaptcha failed", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 2. 生成验证码
	err := logic.GenerateAndSendVerifyCode(req.Email)
	if err != nil {
		zap.L().Error("logic.GenerateAndSendVerifyCode failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 3. 返回响应
	ctx.JSON(http.StatusOK, gin.H{"message": "验证码发送成功"})
}
