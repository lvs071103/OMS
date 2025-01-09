package logic

import (
	"github.com/oms/dao/mysql"
	"github.com/oms/models"
)

func ServerList(page, pageSize int64) (data *models.RespServersList, err error) {

	// 查询服务器列表
	data, err = mysql.ServerList(page, pageSize)
	if err != nil {
		return nil, err
	}

	total, err := mysql.ServersCount()
	if err != nil {
		return nil, err
	}
	data.Total = total
	return
}
