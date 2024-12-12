package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oms/controller"
	_ "github.com/oms/docs" // 这里需要替换为你的项目路径
	"github.com/oms/logger"
	"github.com/oms/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title OMS API
// @version 1.0
// @description OMS是一套运维操作管理系统
// @termsOfService http://localhost:8000/terms/

// @contact.name API Support
// @contact.url http://localhost:8000/support
// @contact.email furong.zhou@dominos.com.cn

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api/v1

func SetupRoute(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 应用 CORS 中间件
	r.Use(middlewares.CORSMiddleware())
	v1 := r.Group("/api/v1")
	// 注册业务
	{
		v1.POST("/signup", controller.SignUpHandler)
		v1.POST("/login", controller.LoginHandler)
		v1.POST("/logout", controller.LogoutHandler) // 退出登录
		v1.POST("/captcha", controller.GenerateCaptchaHandler)
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// 应用JWT中间件
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		// 权限
		v1.GET("/permission/list", controller.PermissionListHander)
		// 用户组
		v1.GET("/group/list", controller.GroupListHander)
		v1.POST("/group/add", controller.GroupAddHandler)
		v1.GET("/group/:id", controller.GroupDetailHandler)
		v1.POST("/group/:id", controller.GroupUpdateHandler)
		v1.DELETE("/group/:id", controller.GroupDeleteHandler)
		// 用户管理
		v1.GET("/user/list", controller.UserListHander)
		v1.GET("/user/:id", controller.UserDetailHandler)
		v1.POST("/user/:id", controller.UserUpdateHandler)
		v1.DELETE("/user/:id", controller.UserDeleteHandler)
		v1.POST("/user/add", controller.UserAddHandler)
		// 环境管理
		v1.GET("/sys/config/env/list", controller.EnvListHandler)
		v1.POST("/sys/config/env/add", controller.EnvAddHandler)
		v1.GET("/sys/config/env/:id", controller.EnvDetailHandler)
		v1.POST("/sys/config/env/:id", controller.EnvUpdateHandler)
		v1.DELETE("/sys/config/env/:id", controller.EnvDeleteHandler)
		// Jenins instance
		v1.GET("/app/release/jenkins/list", controller.JenkinsInstancesListHandler)
		v1.POST("/app/release/jenkins/add", controller.JenkinsInstanceAddHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})

	return r
}
