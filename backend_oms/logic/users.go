package logic

import (
	"github.com/oms/dao/mysql"
	"github.com/oms/models"
	"go.uber.org/zap"
)

// UserList - 用户列表
func UserList(page, pageSize int64) (data *models.RespAuthUserList, err error) {
	data, err = mysql.UserList(page, pageSize)
	if err != nil {
		zap.L().Error("mysql.UserList failed", zap.Error(err))
		return
	}
	// data = users
	return
}

// UserDetail - 用户详情
func UserDetail(uid int64) (data *models.RespAuthUser, err error) {
	data, err = mysql.UserDetail(uid)
	if err != nil {
		zap.L().Error("query user detail failed", zap.Error(err))
		return
	}
	return
}

// UserUpdate - 用户更新
func UserUpdate(id int64, user *models.CreateUserRequest) (err error) {
	err = mysql.UserUpdate(id, user)
	if err != nil {
		zap.L().Error("mysql UserUpdate failed", zap.Error(err))
		return err
	}
	return
}

// UserAdd - 用户添加
func UserAdd(user *models.CreateUserRequest) (err error) {
	// 添加用户
	err = mysql.UserAdd(user)
	if err != nil {
		zap.L().Error("mysql.UserAdd failed", zap.Error(err))
		return err
	}
	return
}

// UserDelete - 删除用户
func UserDelete(id int64) (err error) {
	err = mysql.UserDelete(id)
	if err != nil {
		zap.L().Error("mysql.UserDelete failed", zap.Error(err))
		return err
	}
	return
}
