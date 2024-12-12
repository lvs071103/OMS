package models

import (
	"time"
)

type ModelContentType struct {
	ID       int64  `gorm:"primarykey;type:bigint(20)"`
	AppLabel string `gorm:"column:app_label;type:varchar(100)"`
	Model    string `gorm:"column:model;type:varchar(100);unique"`
}

type AuthPermission struct {
	ID            int64  `gorm:"primarykey;bigint(20)"`
	Name          string `gorm:"column:name;type:varchar(255)"`
	ContentTypeID int64  `gorm:"column:content_type_id;type:bigint(20)"`
	CodeName      string `gorm:"column:codename;type:varchar(100);unique"`
}

type AuthGroup struct {
	ID   int64  `gorm:"column:id;primaryKey;type:bigint(20)"`
	Name string `gorm:"column:name;type:varchar(128);"`
	Desc string `gorm:"column:desc;type:text;"`
}

type AuthGroupPermissions struct {
	ID           int64 `gorm:"primaryKey;bigint(20)"`
	GroupID      int64 `gorm:"column:group_id"`
	PermissionID int64 `gorm:"column:permission_id"`
}

// 自定义关联表结构
type AuthUserGroup struct {
	ID      int64 `gorm:"primaryKey;bigint(20)"`
	UserID  int64 `gorm:"column:user_id"`
	GroupID int64 `gorm:"column:group_id"`
}

// 结构体，对应数据库中的user表
type AuthUser struct {
	ID          int64     `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	UserName    string    `gorm:"column:username;type:varchar(150);unique;not null"`
	Password    string    `gorm:"column:password;type:varchar(128);not null"`
	LastLogin   time.Time `gorm:"column:last_login;type:datetime(6)"`
	Email       string    `gorm:"column:email;type:varchar(254);not null"`
	DateJoined  time.Time `gorm:"column:date_joined;type:datetime(6);not null"`
	FirstName   string    `gorm:"column:first_name;type:varchar(150);not null;default:''"`
	LastName    string    `gorm:"column:last_name;type:varchar(150);not null;default:''"`
	IsSuperuser bool      `gorm:"column:is_superuser;type:tinyint(1);not null;default:false"`
	IsStaff     bool      `gorm:"column:is_staff;type:tinyint(1);not null;default:false"`
	IsActive    bool      `gorm:"column:is_active;type:tinyint(1);not null;default:false"`
	Age         int       `gorm:"column:age;type:tinyint(3);not null;default:0"`
	Gender      int       `gorm:"column:gender;type:tinyint(1);not null;default:0"`
	Job         string    `gorm:"column:job;type:varchar(150);not null;default:''"`
	Address     string    `gorm:"column:address;type:varchar(254);not null;default:''"`
	Desc        string    `gorm:"column:desc;type:text"`
}

// AuthUserPermissions - 用户权限关联表
type AuthUserPermissions struct {
	ID           int64 `gorm:"primaryKey;bigint(20)"`
	UserID       int64 `gorm:"column:user_id"`
	PermissionID int64 `gorm:"column:permission_id"`
}
