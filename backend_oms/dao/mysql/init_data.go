package mysql

import (
	"github.com/oms/models"
	"github.com/oms/pkg/snowflake"
	"go.uber.org/zap"
)

func CheckPermissionExists(codename string) (id int64, err error) {
	sql := `select id from auth_permissions where codename = ?`
	err = db.Get(&id, sql, codename)
	return id, err
}

func CheckModeContentTypeExists(appLabel, model string) (id int64, err error) {
	sql := `select id from model_content_types where app_label = ? and model = ?`
	err = db.Get(&id, sql, appLabel, model)
	return id, err
}

func CommonInsertProcess(perms []models.AuthPermissionRequest, v models.RespModelContentType) (err error) {
	for _, p := range perms {
		// 查询权限是否存在
		id, err := CheckPermissionExists(p.CodeName)
		// 如果有错误，直接返回
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				p.ID = snowflake.GenID()
				sql := `insert into auth_permissions (id, name, codename, content_type_id) values(?,?,?,?)`
				_, err := db.Exec(sql, p.ID, p.Name, p.CodeName, v.ID)
				if err != nil {
					zap.L().Error("insert into auth_permissions failed", zap.Error(err))
					return err
				}
			} else {
				zap.L().Error("CheckPermissionExists failed", zap.Error(err))
				return err
			}
		}
		// 如果id大于0，说明存在，直接跳过
		if id > 0 {
			return nil
		}
	}
	return nil
}

func InitData() {
	// 1. 插入ContentTypes数据
	for _, v := range InitContentTypesData {
		// 查询是否存在
		id, err := CheckModeContentTypeExists(v.AppLabel, v.Model)
		// 如果有错误
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				v.ID = snowflake.GenID()
				sql := `insert into model_content_types (id, app_label, model) values(?,?,?)`
				_, err := db.Exec(sql, v.ID, v.AppLabel, v.Model)
				if err != nil {
					zap.L().Error("insert into model_content_types failed", zap.Error(err))
					return
				}
			} else {
				zap.L().Error("CheckModeContentTypeExists failed", zap.Error(err))
				return
			}
		}
		// 如果id大于0，说明存在，直接跳过
		if id > 0 {
			continue
		}
	}
	zap.L().Info("初始化内容类型数据成功!")

	// 3. 初始化权限数据
	sql := `Select * from model_content_types`
	var contentTypes []models.RespModelContentType
	err := db.Select(&contentTypes, sql)

	if err != nil {
		zap.L().Error("db.Select failed", zap.Error(err))
		return
	}

	for _, v := range contentTypes {
		if v.Model == "logentry" {
			var permissions = InitLogEntry
			err = CommonInsertProcess(permissions, v)
			if err != nil {
				zap.L().Error("CommonInsertProcess failed", zap.Error(err))
				return
			}
		}

		if v.Model == "group" {
			var permissions = InitGroups
			err = CommonInsertProcess(permissions, v)
			if err != nil {
				zap.L().Error("CommonInsertProcess failed", zap.Error(err))
				return
			}
		}
		if v.Model == "permission" {
			var permissions = InitPermissions
			err = CommonInsertProcess(permissions, v)
			if err != nil {
				zap.L().Error("CommonInsertProcess failed", zap.Error(err))
				return
			}
		}
		if v.Model == "user" {
			var permissions = InitUser
			err = CommonInsertProcess(permissions, v)
			if err != nil {
				zap.L().Error("CommonInsertProcess failed", zap.Error(err))
				return
			}
		}

		if v.Model == "contenttype" {
			var permissions = InitContentType
			err = CommonInsertProcess(permissions, v)
			if err != nil {
				zap.L().Error("CommonInsertProcess failed", zap.Error(err))
				return
			}
		}
		if v.Model == "session" {
			var permissions = InitSession
			err = CommonInsertProcess(permissions, v)
			if err != nil {
				zap.L().Error("CommonInsertProcess failed", zap.Error(err))
				return
			}
		}

		if v.Model == "env" {
			var permissions = InitEnv // 这里的权限是自定义的
			err = CommonInsertProcess(permissions, v)
			if err != nil {
				zap.L().Error("CommonInsertProcess failed", zap.Error(err))
				return
			}
		}
	}
	zap.L().Info("初始化权限数据成功!")
}
