package logic

import (
	"github.com/oms/dao/mysql"
	"github.com/oms/models"
	"go.uber.org/zap"
)

// JenkinsInstancesList - Jenkins实例列表
func JenkinsInstancesList(page, page_size int64) (data *models.RespJenkinsInstancesList, err error) {
	// 查询Jenkins实例列表
	data, err = mysql.JenkinsInstancesList(page, page_size)
	if err != nil {
		zap.L().Error("logic JenkinsInstancesList failed", zap.Error(err))
		return
	}

	// 统计Jenkins实例数量
	total, err := mysql.JenkinsInstancesCount()
	if err != nil {
		return nil, err
	}
	data.Total = total

	return

}

// JenkinsInstanceAdd - 添加Jenkins实例
func JenkinsInstanceAdd(req *models.CreateJenkinsInstanceRequest) (err error) {
	// 检查Jenkins实例是否存在
	err = mysql.JenkinsInstanceCheck(req)
	if err != nil {
		zap.L().Error("logic JenkinsInstanceCheck failed", zap.Error(err))
		return
	}
	err = mysql.JenkinsInstanceAdd(req)
	if err != nil {
		zap.L().Error("logic JenkinsInstanceAdd failed", zap.Error(err))
		return
	}
	return
}

// JenkinsInstanceDetail - Jenkins实例详情
func JenkinsInstanceDetail(id int64) (data *models.RespJenkinsInstanceDetail, err error) {
	// 查询Jenkins实例详情
	data, err = mysql.JenkinsInstanceDetail(id)
	if err != nil {
		zap.L().Error("logic JenkinsInstanceDetail failed", zap.Error(err))
		return
	}

	return data, nil
}

// JenkinsInstanceUpdate - 更新Jenkins实例
func JenkinsInstanceUpdate(id int64, req *models.CreateJenkinsInstanceRequest) (err error) {
	// 更新Jenkins实例
	err = mysql.JenkinsInstanceUpdate(id, req)
	if err != nil {
		zap.L().Error("logic JenkinsInstanceUpdate failed", zap.Error(err))
		return
	}
	return
}

// JenkinsInstanceDelete - 删除Jenkins实例
func JenkinsInstanceDelete(id int64) (err error) {
	// 删除Jenkins实例
	err = mysql.JenkinsInstanceDelete(id)
	if err != nil {
		zap.L().Error("logic JenkinsInstanceDelete failed", zap.Error(err))
		return
	}
	return
}

// Release Jobs List
func ReleaseJobsList(page, pageSize int64) (data *models.RespReleaseJobsList, err error) {
	// 查询发布任务列表
	data, err = mysql.ReleaseJobsList(page, pageSize)
	if err != nil {
		zap.L().Error("logic ReleaseJobsList failed", zap.Error(err))
		return
	}

	// 统计发布任务数量
	total, err := mysql.ReleaseJobsCount()
	if err != nil {
		return nil, err
	}
	data.Total = total

	return
}
