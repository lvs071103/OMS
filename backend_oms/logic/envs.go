package logic

import (
	"github.com/oms/dao/mysql"
	"github.com/oms/models"
	"go.uber.org/zap"
)

func EnvList(page, pageSize int64) (data *models.RespOmsEnvList, err error) {
	data, err = mysql.EnvList(page, pageSize)
	if err != nil {
		zap.L().Error("mysql.EnvList failed", zap.Error(err))
		return
	}
	return
}

func EnvAddHandler(env *models.CreateEnvRequest) (err error) {
	err = mysql.EnvAdd(env)
	if err != nil {
		zap.L().Error("mysql.EnvAdd failed", zap.Error(err))
		return
	}
	return
}

// EnvDetail - 环境详情
func EnvDetail(id int64) (data *models.RespOmsEnv, err error) {
	data, err = mysql.EnvDetail(id)
	return
}

// EnvUpdate - 更新环境
func EnvUpdate(id int64, env *models.CreateEnvRequest) (err error) {
	err = mysql.EnvUpdate(id, env)
	if err != nil {
		zap.L().Error("logic EnvUpdate failed", zap.Error(err))
		return
	}
	return
}

// EnvDelete - 删除环境
func EnvDelete(id int64) (err error) {
	err = mysql.EnvDelete(id)
	if err != nil {
		zap.L().Error("logic EnvDelete failed", zap.Error(err))
		return
	}
	return
}
