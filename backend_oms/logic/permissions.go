package logic

import (
	"github.com/oms/dao/mysql"
	"github.com/oms/models"
	"go.uber.org/zap"
)

// PermissionList - 权限列表
func PermissionList() (data []*models.RespPermissionLit, err error) {
	permissions, err := mysql.PermissionList()
	if err != nil {
		zap.L().Error("mysql.PermissionList failed", zap.Error(err))
		return
	}
	data = permissions
	return
}

// PermissionAdd - 添加权限
func PermissionAdd(permission *models.CreatePermissionRequest) (err error) {
	err = mysql.PermissionAdd(permission)
	if err != nil {
		zap.L().Error("mysql.PermissionAdd failed", zap.Error(err))
		return
	}
	return
}
