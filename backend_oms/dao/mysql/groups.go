package mysql

import (
	"database/sql"
	"fmt"

	"github.com/oms/models"
	"github.com/oms/pkg/snowflake"
	"go.uber.org/zap"
)

// 检查用户组是不是已经存在
func CheckGroupIsExists(groupName string) (err error) {
	sqlStr := `select count(id) from auth_groups where name = ?`
	var count int
	if err = db.Get(&count, sqlStr, groupName); err != nil {
		return
	}
	if count > 0 {
		return ErrorGroupExist
	}
	return
}

// GroupCount - 统计用户组数量
func GroupCount() (count int64, err error) {
	sql := `SELECT COUNT(id) FROM auth_groups`
	err = db.Get(&count, sql)
	if err != nil {
		zap.L().Error("query group count failed", zap.Error(err))
	}
	return
}

// GroupList - 用户组列表
func GroupList(page, pageSize int64) (groups []models.RespAuthGroup, err error) {
	sql := `SELECT id, name FROM auth_groups ORDER BY id LIMIT ?,?`
	// make([]T, len, cap) 生成一个切片，长度为len，容量为cap
	groups = make([]models.RespAuthGroup, 0, pageSize)
	err = db.Select(&groups, sql, (page-1)*pageSize, pageSize)
	if err != nil {
		zap.L().Error("db.Select failed", zap.Error(err))
		return
	}

	return groups, nil
}

// GetGroupPermission - 查询用户组权限
func GetGroupPermission(groupID int64) (permissions []models.RespPermission, err error) {
	sql := `SELECT p.id, p.name, p.codename, p.content_type_id, 
	c.id AS "content_types.id", 
	c.model AS "content_types.model"
	FROM auth_group_permissions gp 
	LEFT JOIN auth_permissions p ON gp.permission_id = p.id 
	JOIN model_content_types c ON c.id = p.content_type_id 
	WHERE gp.group_id = ?;`
	err = db.Select(&permissions, sql, groupID)
	if err != nil {
		zap.L().Error("query permissions list failed", zap.Error(err))
	}
	return
}

func GroupDetail(id int64) (data *models.RespAuthGroup, err error) {
	data = new(models.RespAuthGroup)
	// 取出用户组数据
	sql1 := "SELECT `id`,`name`,`desc` FROM auth_groups where id = ?"
	if err = db.Get(data, sql1, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Error("query group failed", zap.Error(err))
			return nil, err
		}
	}

	return
}

// AddGroup - 添加用户组和权限
func AddGroup(group *models.CreateGroupRequest) (err error) {
	// 检测用户组是否存在
	err = CheckGroupIsExists(group.Name)
	if err != nil {
		return
	}
	// 生成雪花ID-GID
	group.ID = snowflake.GenID()
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		return
	}
	// 插入用户组
	sql := "INSERT INTO auth_groups (`id`, `name`, `desc`) VALUES (?, ?, ?);"
	_, err = tx.Exec(sql, group.ID, group.Name, group.Desc)
	if err != nil {
		tx.Rollback()
		return
	}
	// 插入用户组权限
	if len(group.PermissionIDS) == 0 {
		// 如果没有权限，直接提交事务
		tx.Commit()
		return
	} else {
		// 如果有权限，插入权限
		for _, pid := range group.PermissionIDS {
			id := snowflake.GenID()
			sql = "INSERT INTO auth_group_permissions (`id`, `group_id`, `permission_id`) VALUES (?, ?, ?)"
			_, err = tx.Exec(sql, id, group.ID, pid)
			if err != nil {
				tx.Rollback()
				return
			}
		}
		tx.Commit()
		return
	}
}

func GroupUpdate(id int64, group *models.CreateGroupRequest) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	sql := "UPDATE auth_groups SET `name` = ?, `desc`=? WHERE `id` = ?"
	_, err = tx.Exec(sql, group.Name, group.Desc, id)
	if err != nil {
		zap.L().Error("db.Exec failed", zap.Error(err))
		tx.Rollback()
		return
	}

	// 删除用户组权限
	sql = "DELETE FROM auth_group_permissions WHERE `group_id` = ?"
	_, err = tx.Exec(sql, id)
	if err != nil {
		zap.L().Error("db.Exec failed", zap.Error(err))
		tx.Rollback()
		return
	}
	// 插入用户组权限
	if len(group.PermissionIDS) == 0 {
		// 如果没有权限，直接提交事务
		tx.Commit()
		return
	} else {
		// 如果有权限，插入权限
		for _, pid := range group.PermissionIDS {
			sql = "INSERT INTO auth_group_permissions (`group_id`, `permission_id`) values (?, ?)"
			_, err = tx.Exec(sql, id, pid)
			if err != nil {
				zap.L().Error("db.Exec failed", zap.Error(err))
				tx.Rollback()
				return
			}
		}
		tx.Commit()
		return
	}
}

// GroupDelete - 删除用户组
func GroupDelete(id int64) (err error) {
	fmt.Println("id:", id)
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		return
	}

	// 删除用户组
	sql := `delete from auth_groups where id = ?`
	_, err = db.Exec(sql, id)
	if err != nil {
		zap.L().Error("db.Exec failed", zap.Error(err))
		tx.Rollback()
		return
	}

	// 删除用户组权限
	sql = `delete from auth_group_permissions where group_id = ?`
	_, err = db.Exec(sql, id)
	if err != nil {
		zap.L().Error("db.Exec failed", zap.Error(err))
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
