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