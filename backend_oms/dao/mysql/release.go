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

func JenkinsInstanceDetail(id int64) (data *models.RespJenkinsInstanceDetail, err error) {
	// 实例化结构体
	data = new(models.RespJenkinsInstanceDetail)
	sql := `SELECT a.id, a.name, a.env_id, a.address, a.auth_type,
	a.username, a.password, a.desc, b.name AS "env_name"
	FROM oms_jenkins_instances a 
	LEFT JOIN oms_env_configs b ON a.env_id = b.id 
	WHERE a.id = ?`
	err = db.Get(data, sql, id)
	if err != nil {
		zap.L().Error("query oms_jenkins_instances failed", zap.Error(err))
		return nil, err
	}

	return
}

func JenkinsInstanceUpdate(id int64, req *models.CreateJenkinsInstanceRequest) (err error) {
	sql := "UPDATE oms_jenkins_instances " +
		"SET `env_id`=?, `name`=?, `address`=?, `auth_type`=?, `username`=?, `password`=?, `desc`=? " +
		"WHERE id=?"
	_, err = db.Exec(sql,
		req.EnvID,
		req.Name,
		req.Address,
		req.AuthType,
		req.UserName,
		req.Password,
		req.Desc,
		id)
	if err != nil {
		zap.L().Error("update oms_jenkins_instances failed", zap.Error(err))
		return
	}
	return
}

func JenkinsInstanceDelete(id int64) (err error) {
	sql := "DELETE FROM oms_jenkins_instances WHERE id = ?"
	_, err = db.Exec(sql, id)
	if err != nil {
		zap.L().Error("delete oms_jenkins_instances failed", zap.Error(err))
		return
	}
	return
}

// ReleaseJobsCount - 发布任务总数
func ReleaseJobsCount() (total int64, err error) {
	total_sql := `SELECT count(0) FROM release_jobs`
	err = db.Get(&total, total_sql)
	if err != nil {
		zap.L().Error("db.Get failed", zap.Error(err))
		return 0, err
	}
	return
}

// ReleaseJobsList - 发布任务列表
func ReleaseJobsList(page, pageSize int64) (data *models.RespReleaseJobsList, err error) {
	// 实例化结构体
	data = new(models.RespReleaseJobsList)
	sql := "SELECT `id`, `name`, `deploy_name`, `servce_name`, `desc` FROM release_jobs LIMIT ?, ?"
	err = db.Select(&data.Jobs, sql, (page-1)*pageSize, pageSize)
	if err != nil {
		zap.L().Error("query release_jobs failed", zap.Error(err))
		return nil, err
	}

	return
}
