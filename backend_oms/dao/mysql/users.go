package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/oms/models"
	"github.com/oms/pkg/snowflake"
	"go.uber.org/zap"
)

// 把每一步数据库操作封装成函数

const secrect_string = "oms"

// CheckUserIsExists 检查指定用户是否存在
func CheckUserIsExists(username string) (err error) {
	// 查询用户
	sql := "select count(id) from auth_users where username = ?"
	var count int
	if err := db.Get(&count, sql, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// CheckUserIsExistsByUid 根据uid判断用户是否存在
func CheckUserIsExistsByUid(id int64) (err error) {
	sql := `select count(0) from auth_users where id = ?`
	var count int
	if err := db.Get(&count, sql, id); err != nil {
		return err
	}
	if count == 0 {
		return ErrorUserNotExist
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.AuthUser) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	user.DateJoined = time.Now()
	user.IsActive = true
	// 执行sql语句入库
	sql := `insert into auth_users 
	(id,username,email,password,date_joined,
	is_active) 
	values(?,?,?,?,?,?)`

	_, err = db.Exec(
		sql, user.ID,
		user.UserName,
		user.Email,
		user.Password,
		user.DateJoined,
		user.IsActive,
	)
	return
}

// encryptPassword 密码加盐
func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secrect_string))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func UpdateLastLogin(user *models.ResponseUser) (err error) {
	user.LastLogin = time.Now()
	fmt.Println(user.LastLogin)
	sql := `update auth_users set last_login = ? where username = ?`
	_, err = db.Exec(sql, user.LastLogin, user.UserName)
	return
}

func Login(user *models.ResponseUser) (err error) {
	// 先保存一下请求的password
	reqPassword := user.Password
	sqlStr := "select id, username, password from auth_users where username=?"
	// 将查询结果存入user
	err = db.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return ErrorInvalidPassword
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正常
	password := encryptPassword(reqPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	// 更新last_login字段
	err = UpdateLastLogin(user)
	if err != nil {
		zap.L().Error("update last_login field failed", zap.Error(err))
		return err
	}
	return
}

// UserDetail 根据用户id获取用户详情
func UserDetail(id int64) (user *models.RespAuthUser, err error) {
	// 1.查询用户信息
	user = new(models.RespAuthUser)
	user_sql := "select * from auth_users where id = ?"
	if err = db.Get(user, user_sql, id); err != nil {
		zap.L().Error("query user detail failed", zap.Error(err))
		return
	}

	// 2.查询用户权限信息
	perms_sql := `SELECT p.id, p.name, p.codename, p.content_type_id, 
	c.id AS "content_types.id", c.model AS "content_types.model"
	FROM auth_user_permissions up 
	LEFT JOIN auth_permissions p ON up.permission_id = p.id 
	JOIN model_content_types c ON c.id = p.content_type_id 
	WHERE up.user_id = ?;`
	var permissions []models.RespPermission
	err = db.Select(&permissions, perms_sql, id)
	if err != nil {
		zap.L().Error("query auth_user_permissions failed", zap.Error(err))
		return nil, err
	}
	user.Permissions = permissions

	// 3.查询用户组信息
	groups_sql := `SELECT g.id, g.name FROM auth_users u 
	INNER JOIN auth_user_groups ug ON u.id = ug.user_id
	INNER JOIN auth_groups g ON ug.group_id = g.id
	WHERE u.id = ?;`
	var groups []models.RespAuthGroup
	err = db.Select(&groups, groups_sql, id)
	if err != nil {
		zap.L().Error("query auth_user_groups failed", zap.Error(err))
		return nil, err
	}
	user.Groups = groups
	return
}

func UserList(page, pageSize int64) (data *models.RespAuthUserList, err error) {
	// 1.初始化结构体
	data = new(models.RespAuthUserList)
	// 2.查询用户列表
	users_sql := `select * from auth_users ORDER BY id limit ?,?`
	err = db.Select(&data.Users, users_sql, (page-1)*pageSize, pageSize)
	if err != nil {
		zap.L().Error("db.Select failed", zap.Error(err))
		return nil, err
	}
	// 2.查询总数
	total_sql := `select count(0) from auth_users`
	err = db.Get(&data.Total, total_sql)
	if err != nil {
		zap.L().Error("db.Get failed", zap.Error(err))
		return nil, err
	}

	return
}

// UserUpdate - 用户更新
func UserUpdate(id int64, user *models.CreateUserRequest) (err error) {
	// 1. 开启事务
	tx, err := db.Begin()
	if err != nil {
		zap.L().Error("begin tx failed", zap.Error(err))
		return
	}
	// 2.更新用户表
	sql := `update auth_users set 
	username = ?,
	is_staff = ?,
	is_superuser = ?,
	is_active = ?,
	email = ?,
	first_name = ?,
	last_name = ?,
	age = ?,
	gender = ?,
	job = ?,
	address = ?
	where id = ?`
	_, err = db.Exec(sql,
		user.UserName,
		user.IsStaff,
		user.IsSuperuser,
		user.IsActive,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Age,
		user.Gender,
		user.Job,
		user.Address,
		id,
	)

	if err != nil {
		zap.L().Error("update auth_users failed", zap.Error(err))
		tx.Rollback()
		return err
	}

	var count int
	// 2. 更新用户和组关联表
	// 检查用户组关联表是否存在user_id
	sql = "select count(0) from auth_user_groups where `user_id` = ?"
	if err := db.Get(&count, sql, id); err != nil {
		tx.Rollback()
		return err
	}
	// 2.1 删除用户和组关联表
	if count == 0 {
		zap.L().Info("user and group relation is not exists")
		// 如果不存在，直接跳过
	} else {
		// 如果存在，删除
		sql = "DELETE from auth_user_groups where `user_id` = ?"
		_, err = tx.Exec(sql, id)
		if err != nil {
			zap.L().Error("delete user and group relation failed", zap.Error(err))
			tx.Rollback()
			return err
		}
	}

	// 2.2 插入用户和组关联表
	if len(user.Groups) != 0 {
		for _, gid := range user.Groups {
			l_id := snowflake.GenID()
			sql = "INSERT INTO auth_user_groups (`id`, `user_id`, `group_id`) VALUES (?, ?, ?)"
			_, err = tx.Exec(sql, l_id, id, gid)
			if err != nil {
				zap.L().Error("insert user and group relation failed", zap.Error(err))
				tx.Rollback()
				return err
			}
		}
	} else {
		// 如果user.group为空，直接跳过
		zap.L().Info("user.group is empty")
	}

	// 3. 更新用户和权限关联表
	// 检查用户权限关联表是否存在user_id
	sql = "select count(0) from auth_user_permissions where `user_id` = ?"
	if err := db.Get(&count, sql, id); err != nil {
		tx.Rollback()
		return err
	}
	// 3.1 删除用户和权限关联表
	if count == 0 {
		// 如果不存在，直接跳过
		zap.L().Info("user and permission relation is not exists")
	} else {
		// 如果存在，删除
		sql = "DELETE from auth_user_permissions where `user_id` = ?"
		_, err = tx.Exec(sql, id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 3.2 插入用户和权限关联表
	if len(user.PermissionIDS) != 0 {
		for _, pid := range user.PermissionIDS {
			l_id := snowflake.GenID()
			sql = "INSERT INTO auth_user_permissions (`id`, `user_id`, `permission_id`) VALUES (?, ?, ?)"
			_, err = tx.Exec(sql, l_id, id, pid)
			if err != nil {
				zap.L().Error("insert user and permission relation failed", zap.Error(err))
				tx.Rollback()
				return err
			}
		}
	} else {
		// 如果user.permission为空，直接跳过
		zap.L().Info("user.permissions is empty")
	}

	tx.Commit()
	return
}

// UserAdd - 添加用户
func UserAdd(user *models.CreateUserRequest) (err error) {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// 1.判断用户是否存在
	if err := CheckUserIsExists(user.UserName); err != nil {
		tx.Rollback()
		return err
	}
	// 2.生成uid
	user.ID = snowflake.GenID()
	user.Password = encryptPassword(user.Password)
	user.DateJoined = time.Now()

	sql := `INSERT INTO auth_users (id, username, email, password,
	date_joined, first_name, last_name, is_superuser,
	is_staff, is_active, age,
	gender, job, address) 
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	_, err = tx.Exec(sql,
		user.ID,
		user.UserName,
		user.Email,
		user.Password,
		user.DateJoined,
		user.FirstName,
		user.LastName,
		user.IsSuperuser,
		user.IsStaff,
		user.IsActive,
		user.Age,
		user.Gender,
		user.Job,
		user.Address,
	)
	if err != nil {
		tx.Rollback()
		return
	}

	// 如果user.group为列表
	if len(user.Groups) != 0 {
		// 遍历插入
		for _, gid := range user.Groups {
			// 插入用户和组关联表
			id := snowflake.GenID()
			sql = `insert into auth_user_groups (id, user_id, group_id) values (?,?,?)`
			_, err = tx.Exec(sql, id, user.ID, gid)
			if err != nil {
				zap.L().Error("insert user and group relation failed", zap.Error(err))
				tx.Rollback()
				return
			}
		}
	} else {
		// 如果user.group为空，直接跳过
		zap.L().Info("user.group is empty")
	}

	// 3.如果没有权限 直接返回
	if len(user.PermissionIDS) == 0 {
		// 跳过
		zap.L().Info("user.permissions is empty")
	} else {
		// 如果有权限，插入权限
		for _, pid := range user.PermissionIDS {
			// 向用户和权限关联表中插入数据
			id := snowflake.GenID()
			sql = "INSERT INTO auth_user_permissions (`id`, `user_id`, `permission_id`) VALUES (?, ?, ?)"
			_, err = tx.Exec(sql, id, user.ID, pid)
			if err != nil {
				zap.L().Error("insert user permission relation failed", zap.Error(err))
				tx.Rollback()
				return
			}
		}
	}

	// 4.提交事务
	tx.Commit()
	return
}

// UserDelete 删除用户
func UserDelete(id int64) (err error) {
	// 1.判断用户是否存在
	if err := CheckUserIsExistsByUid(id); err != nil {
		return err
	}
	// 2.开启事务
	tx, err := db.Begin()
	if err != nil {
		return
	}
	// 3.删除用户
	sql := "DELETE from auth_users where id = ?"
	_, err = tx.Exec(sql, id)
	if err != nil {
		zap.L().Error("delete user failed", zap.Error(err))
		tx.Rollback()
		return err
	}
	// 4. 删除用户和组关联表
	sql = "DELETE from auth_user_groups where user_id = ?"
	_, err = tx.Exec(sql, id)
	if err != nil {
		zap.L().Error("delete user and group relation failed", zap.Error(err))
		tx.Rollback()
		return err
	}
	// 5. 删除用户和权限关联表
	sql = "DELETE from auth_user_permissions where user_id = ?"
	_, err = tx.Exec(sql, id)
	if err != nil {
		zap.L().Error("delete user and permission relation failed", zap.Error(err))
		tx.Rollback()
		return err
	}
	// 6. 提交事务
	tx.Commit()
	return
}
