package logic

import (
	"strconv"

	"github.com/oms/dao/mysql"
	"github.com/oms/models"
	"go.uber.org/zap"
)

// GroupList - 用户组列表
func GroupList(page, pageSize int64) (data *models.RespGroupList, err error) {
	// 查询用户组列表
	groups, err := mysql.GroupList(page, pageSize)
	if err != nil {
		zap.L().Error("mysql.GroupList failed", zap.Error(err))
		return
	}

	// 查询每个组的权限
	for i := range groups {
		id, _ := strconv.ParseInt(groups[i].ID, 10, 64)
		permissions, err := mysql.GetGroupPermission(id)
		if err != nil {
			zap.L().Error("get group permission failed", zap.Error(err))
			return nil, err
		}
		groups[i].Permissions = permissions
	}

	// 统计用户组数量
	total, err := mysql.GroupCount()
	if err != nil {
		return nil, err
	}
	data = new(models.RespGroupList)
	data.Groups = groups
	data.Total = total
	return
}

// AddGroupLogic - 添加用户组
func AddGroupLogic(group *models.CreateGroupRequest) (err error) {
	err = mysql.AddGroup(group)
	if err != nil {
		zap.L().Error("mysql.AddGroup failed", zap.Error(err))
		return
	}
	return
}

// GroupDetail - 用户组详情
func GroupDetail(id int64) (groups *models.RespAuthGroup, err error) {
	// 查询用户组详情
	groups, _ = mysql.GroupDetail(id)
	// 查询用户组权限
	permissions, _ := mysql.GetGroupPermission(id)
	groups.Permissions = permissions
	return
}

// GroupUpdate - 更新用户组
func GroupUpdate(id int64, group *models.CreateGroupRequest) (err error) {
	err = mysql.GroupUpdate(id, group)
	if err != nil {
		zap.L().Error("mysql.GroupUpdate failed", zap.Error(err))
		return
	}
	return
}

// GroupDelete - 删除用户组
func GroupDelete(id int64) (err error) {
	err = mysql.GroupDelete(id)
	if err != nil {
		zap.L().Error("mysql.GroupDelete failed", zap.Error(err))
		return
	}
	return
}
