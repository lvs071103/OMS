package mysql

import (
	"github.com/oms/models"
	"github.com/oms/pkg/snowflake"
	"go.uber.org/zap"
)

func CheckEnvIsExists(name string) (err error) {
	sql := `select count(1) from oms_env_configs where name = ?`
	var count int
	err = db.Get(&count, sql, name)
	if err != nil {
		return
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// EnvList - 环境列表
func EnvList(page, pageSize int64) (data *models.RespOmsEnvList, err error) {
	data = new(models.RespOmsEnvList)
	// 如果page、pageSize为0，则查询所有
	if page == 0 || pageSize == 0 {
		sql := `select * from oms_env_configs`
		err = db.Select(&data.Envs, sql)
		if err != nil {
			zap.L().Error("query oms_env_configs failed", zap.Error(err))
			return nil, err
		}
		return
	} else {
		// 1.查询列表
		sql := `select * from oms_env_configs limit ?, ?`
		err = db.Select(&data.Envs, sql, (page-1)*pageSize, pageSize)
		if err != nil {
			zap.L().Error("query oms_env_configs failed", zap.Error(err))
			return nil, err
		}
	}

	// 2.查询总数
	total_sql := `select count(0) from oms_env_configs`
	err = db.Get(&data.Total, total_sql)
	if err != nil {
		zap.L().Error("db.Get failed", zap.Error(err))
		return nil, err
	}

	return
}

// EnvAdd - 添加环境
func EnvAdd(env *models.CreateEnvRequest) (err error) {
	// 检测环境是否存在
	err = CheckEnvIsExists(env.Name)
	if err != nil {
		zap.L().Error("CheckEnvIsExists failed", zap.Error(err))
		return
	}
	// 生成雪花ID
	env.ID = snowflake.GenID()
	// 执行sql语句入库
	sql := "INSERT INTO oms_env_configs (`id`, `name`, `label`, `desc`) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(sql, env.ID, env.Name, env.Label, env.Desc)
	return
}

// EnvDetail - 环境详情
func EnvDetail(id int64) (data *models.RespOmsEnv, err error) {
	data = new(models.RespOmsEnv)
	sql := "select `id`,`name`,`label`,`desc` from oms_env_configs where id = ?"
	err = db.Get(data, sql, id)
	if err != nil {
		zap.L().Error("query oms_env_configs failed", zap.Error(err))
		return nil, err
	}
	return data, nil
}

// EnvUpdate - 更新环境
func EnvUpdate(id int64, env *models.CreateEnvRequest) (err error) {
	sql := " UPDATE oms_env_configs SET `name` = ?, `label` = ?, `desc` = ? WHERE `id` = ?"
	_, err = db.Exec(sql, env.Name, env.Label, env.Desc, id)
	if err != nil {
		zap.L().Error("update oms_env_configs failed", zap.Error(err))
		return
	}
	return
}

// EnvDelete - 删除环境
func EnvDelete(id int64) (err error) {
	sql := "DELETE FROM oms_env_configs WHERE `id` = ?"
	_, err = db.Exec(sql, id)
	if err != nil {
		zap.L().Error("delete oms_env_configs failed", zap.Error(err))
		return
	}
	return
}
