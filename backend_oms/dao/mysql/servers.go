package mysql

import (
	"github.com/oms/models"
	"go.uber.org/zap"
)

func ServersCount() (count int64, err error) {
	sql := "SELECT COUNT(*) FROM oms_servers"
	err = db.Get(&count, sql)
	if err != nil {
		zap.L().Error("query oms_servers failed", zap.Error(err))
		return 0, err
	}

	return
}

// ServerList - 服务器列表
func ServerList(page, pageSize int64) (data *models.RespServersList, err error) {
	// 查询服务器列表
	data = new(models.RespServersList) // 实例化结构体
	sql := "SELECT `id`, `name`, `ip`, `env_id`, `env_name`, `desc` FROM oms_servers LIMIT ?, ?"
	err = db.Select(&data.Servers, sql, (page-1)*pageSize, pageSize)
	if err != nil {
		zap.L().Error("query oms_servers failed", zap.Error(err))
		return nil, err
	}

	return
}
