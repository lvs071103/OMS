package mysql

import (
	"errors"

	"github.com/oms/models"
	"github.com/oms/pkg/snowflake"
	"go.uber.org/zap"
)

func JenkinsInstanceCheck(req *models.CreateJenkinsInstanceRequest) (err error) {
	sql := `SELECT count(id) FROM oms_jenkins_instances WHERE name = ?`
	var count int
	if err := db.Get(&count, sql, req.Name); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("Jenkins实例已存在")
	}
	return
}

func JenkinsInstancesList(page, pageSize int64) (data *models.RespJenkinsInstancesList, err error) {
	// 实例化结构体
	data = new(models.RespJenkinsInstancesList)
	sql := "SELECT `id`, `name`, `address`, `auth_type`, `desc` FROM oms_jenkins_instances LIMIT ?, ?"
	err = db.Select(&data.Instances, sql, (page-1)*pageSize, pageSize)
	if err != nil {
		zap.L().Error("query oms_jenkins_instances failed", zap.Error(err))
		return nil, err
	}

	return
}

func JenkinsInstancesCount() (total int64, err error) {
	total_sql := `SELECT count(0) FROM oms_jenkins_instances`
	err = db.Get(&total, total_sql)
	if err != nil {
		zap.L().Error("db.Get failed", zap.Error(err))
		return 0, err
	}
	return
}

func JenkinsInstanceAdd(req *models.CreateJenkinsInstanceRequest) (err error) {
	id := snowflake.GenID()
	sql := "INSERT INTO oms_jenkins_instances " +
		"(`id`, `env_id`, `name`, `address`, `auth_type`, `username`, `password`, `desc`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(
		sql,
		id,
		req.EnvID,
		req.Name,
		req.Address,
		req.AuthType,
		req.UserName,
		req.Password,
		req.Desc)

	return
}
