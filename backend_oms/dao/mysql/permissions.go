package mysql

import (
	"github.com/oms/models"
	"go.uber.org/zap"
)

func PermissionList() (data []*models.RespPermissionLit, err error) {
	// 关联查询
	sql := `SELECT p.id, p.name, p.content_type_id, p.codename, c.id 
	AS "content_types.id", c.model AS "content_types.model" FROM auth_permissions p
	JOIN model_content_types c ON p.content_type_id = c.id`
	err = db.Select(&data, sql)
	if err != nil {
		zap.L().Error("db.Select failed", zap.Error(err))
		return
	}
	return
}

// PermissionAdd - 添加权限
func PermissionAdd(permission *models.CreatePermissionRequest) (err error) {
	sql := `insert into permissions (name, code_name) values (?, ?)`
	_, err = db.Exec(sql, permission.Name, permission.CodeName)
	if err != nil {
		zap.L().Error("db.Exec failed", zap.Error(err))
		return
	}
	return
}

// PermissionDelete - 删除权限
func PermissionDelete(pid int64) (err error) {
	sql := `delete from permissions where pid = ?`
	_, err = db.Exec(sql, pid)
	if err != nil {
		zap.L().Error("db.Exec failed", zap.Error(err))
		return
	}
	return
}

// PermissionUpdate - 更新权限
func PermissionUpdate(pid int64, permission *models.AuthPermission) (err error) {
	sql := `update permissions set name = ?, code_name = ? where pid = ?`
	_, err = db.Exec(sql, permission.Name, permission.CodeName, pid)
	if err != nil {
		zap.L().Error("db.Exec failed", zap.Error(err))
		return
	}
	return
}

// PermissionDetail - 权限详情
func PermissionDetail(pid int64) (data *models.AuthPermission, err error) {
	data = new(models.AuthPermission)
	sql := `select * from permissions where pid = ?`
	err = db.Get(data, sql, pid)
	if err != nil {
		zap.L().Error("db.Get failed", zap.Error(err))
		return
	}
	return
}
