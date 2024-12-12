package mysql

import (
	"github.com/oms/models"
	"go.uber.org/zap"
)

func Makemigrations() {
	// 方法用于设置多对多关系的连接表。你可以使用这个方法来指定连接表的模型和字段。
	// global.DB.SetupJoinTable(&models.User{}, "Group", &models.Group{})
	// Set()方法用于设置 GORM 的配置选项或上下文值
	err := gdb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.AuthUser{},
		&models.AuthGroup{},
		&models.ModelContentType{},
		&models.AuthPermission{},
		&models.AuthUserGroup{},
		&models.AuthGroupPermissions{},
		&models.AuthUserPermissions{},
		&models.OmsEnvConfig{},
		&models.OmsJenkinsInstances{},
	)
	if err != nil {
		zap.L().Error("生成数据库表结构失败", zap.Error(err))
		return
	}
	zap.L().Info("生成数据库表结构成功!")
}
