package models

import "time"

// CreateGroupRequest 创建用户组参数
type CreateGroupRequest struct {
	ID            int64    `json:"id"`
	Name          string   `json:"name" binding:"required"`
	Desc          string   `json:"desc"`
	PermissionIDS []string `json:"permissions"`
}

// CreateUserRequest 创建用户参数
type CreateUserRequest struct {
	ID            int64     `json:"id"`
	UserName      string    `json:"username" binding:"required"` // binding:"required"表示必填字段
	Password      string    `json:"password" binding:"required"`
	Email         string    `json:"email" binding:"required,email"`
	LastLogin     time.Time `json:"last_login"`
	DateJoined    time.Time `json:"date_joined"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	IsSuperuser   bool      `json:"is_superuser"`
	IsStaff       bool      `json:"is_staff"`
	IsActive      bool      `json:"is_active"`
	Age           int       `json:"age"`
	Gender        int       `json:"gender"`
	Job           string    `json:"job"`
	Address       string    `json:"address"`
	Groups        []string  `json:"groups"`
	PermissionIDS []string  `json:"permissions"`
}

// CreateContentTypeRequest 创建权限
type CreatePermissionRequest struct {
	Name     string `json:"name" binding:"required"`
	CodeName string `json:"code_name" binding:"required"`
}

// CreateEnvRequest 创建环境参数
type CreateEnvRequest struct {
	ID    int64  `json:"id"`
	Name  string `json:"name" binding:"required"`
	Label string `json:"label"`
	Desc  string `json:"desc"`
}

// CaptchaRequest 生成验证码参数
type CaptchaRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ModelContentTypeRequest struct {
	ID       int64  `json:"id"`
	AppLabel string `json:"app_label" binding:"required"`
	Model    string `json:"model" binding:"required"`
}

// AuthPermissionRequest 权限请求参数
type AuthPermissionRequest struct {
	ID            int64  `gorm:"primarykey;bigint(20)"`
	Name          string `gorm:"column:name;type:varchar(255)"`
	ContentTypeID int64  `gorm:"column:content_type_id;type:int(11)"`
	CodeName      string `gorm:"column:codename;type:varchar(100);unique"`
}

// CreateOmsEnvRequest 创建环境请求参数
type CreateOmsEnvRequest struct {
	ID    int64  `json:"id"`
	Name  string `json:"name" binding:"required"`
	Label string `json:"label"`
	Desc  string `json:"desc"`
}

// CreateJenkinsJobRequest 创建jenkins任务请求参数
type CreateJenkinsInstanceRequest struct {
	ID       int64  `json:"id"`
	EnvID    string `json:"env_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	AuthType int    `json:"auth_type"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Desc     string `json:"desc"`
}
