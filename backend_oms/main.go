package main

import (
	"fmt"
	"os"

	"github.com/oms/controller"
	"github.com/oms/dao/mysql"
	"github.com/oms/dao/redis"
	"github.com/oms/logger"
	"github.com/oms/pkg/snowflake"
	"github.com/oms/routes"
	"github.com/oms/settings"
	"go.uber.org/zap"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: oms config.yaml")
		return
	}
	// 加载配置
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	// 日志初始化
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init log failed, err:%v\n", err)
	}
	defer zap.L().Sync()

	// 初始化mysql
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql faile, err:%v/n", err)
		return
	}
	// 延迟关闭
	defer mysql.Close()

	// gorm 初始化
	mysql.InitDB(settings.Conf.MySQLConfig)

	// 延迟关闭
	defer mysql.CloseDB()

	// 是否自动创建表结构
	if settings.Conf.AutoMigration {
		mysql.Makemigrations()
	}

	// 初始化redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	// 初始化snowflake
	if err := snowflake.Init("2024-08-03", 1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 初始化权限数据和ContentType数据
	mysql.InitData()

	// 初始化gin框架内置的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}

	// 注册路由
	r := routes.SetupRoute(settings.Conf.Mode)
	if err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
