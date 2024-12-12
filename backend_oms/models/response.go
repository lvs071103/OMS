package models

import (
	"database/sql"
	"time"
)

// ResponseUser 登陆接口返回的用户信息
type ResponseUser struct {
	ID        string    `db:"id"`
	UserName  string    `db:"username"`
	Password  string    `db:"password"`
	LastLogin time.Time `db:"last_login"`
	Token     string
}

// RespAuthUser 接口返回的用户信息
type RespAuthUser struct {
	ID          string           `db:"id" json:"id"` // json 用于序列化, 如果不写json, 序列化后的字段名为大写
	UserName    string           `db:"username" json:"username"`
	Password    string           `db:"password" json:"password"`
	LastLogin   *time.Time       `db:"last_login" json:"last_login"` // 允许为空
	Email       string           `db:"email" json:"email"`
	DateJoined  time.Time        `db:"date_joined" json:"date_joined"`
	FirstName   string           `db:"first_name" json:"first_name"`
	LastName    string           `db:"last_name" json:"last_name"`
	IsSuperuser bool             `db:"is_superuser" json:"is_superuser"`
	IsStaff     bool             `db:"is_staff" json:"is_staff"`
	IsActive    bool             `db:"is_active" json:"is_active"`
	Age         int              `db:"age" json:"age"`
	Gender      int              `db:"gender" json:"gender"`
	Job         string           `db:"job" json:"job"`
	Address     string           `db:"address" json:"address"`
	Desc        sql.NullString   `db:"desc" json:"desc"`
	Groups      []RespAuthGroup  `json:"groups"`
	Permissions []RespPermission `json:"permissions"`
}

type RespAuthUserList struct {
	Total int            `json:"total"`
	Users []RespAuthUser `json:"users"`
}

// RespContentType 内容类型响应
type RespModelContentType struct {
	ID       string `db:"id" json:"id"`
	AppLabel string `db:"app_label" json:"app_label"`
	Model    string `db:"model" json:"model"`
}

// RespPermissions 权限列表响应
type RespPermissionLit struct {
	ID               string               `db:"id" json:"id"`
	Name             string               `db:"name" json:"name"`
	ContentTypeID    string               `db:"content_type_id" json:"content_type_id"`
	CodeName         string               `db:"codename" json:"codename"`
	ModelContentType RespModelContentType `db:"content_types" json:"content_types"` // 嵌套的内容类型
}

// RespPermission 用户组权限响应
type RespPermission struct {
	ID               string               `db:"id" json:"id"`
	Name             string               `db:"name" json:"name"`
	CodeName         string               `db:"codename" json:"codename"`
	ContentTypeID    string               `db:"content_type_id" json:"content_type_id"`
	ModelContentType RespModelContentType `db:"content_types" json:"content_types"`
}

// RespGroup 用户组响应
type RespAuthGroup struct {
	ID          string           `db:"id" json:"id"`
	Name        string           `db:"name" json:"name"`
	NickName    string           `db:"nick_name" json:"nick_name"`
	Label       string           `db:"label" json:"label"`
	Desc        string           `db:"desc" json:"desc"`
	Permissions []RespPermission `json:"permissions"`
}

// RespGroupList 用户组列表响应
type RespGroupList struct {
	Total  int64           `json:"total"`
	Groups []RespAuthGroup `json:"groups"`
}

// RespOmsEnv 环境响应
type RespOmsEnv struct {
	ID    string `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Label string `db:"label" json:"label"`
	Desc  string `db:"desc" json:"desc"`
}

type RespOmsEnvList struct {
	Total int64        `json:"total"`
	Envs  []RespOmsEnv `json:"envs"`
}

// RespJenkinsInstance jenkins实例响应
type RespJenkinsInstances struct {
	ID       string `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Address  string `db:"address" json:"address"`
	AuthType bool   `db:"auth_type" json:"auth_type"`
	Desc     string `db:"desc" json:"desc"`
}

// RespJenkinsInstancesList jenkins实例列表响应
type RespJenkinsInstancesList struct {
	Total     int64                  `json:"total"`
	Instances []RespJenkinsInstances `json:"instances"`
}
