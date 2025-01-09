package models

import "time"

// OmsJenkinsJobs - jenkins任务定义
type ReleaseJobs struct {
	ID         int64  `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	Name       string `gorm:"column:name;type:varchar(150);unique;not null"`
	DeployName string `gorm:"column:deploy_name;type:varchar(150);not null"`
	ServceName string `gorm:"column:servce_name;type:varchar(150);not null"`
	Desc       string `gorm:"column:desc;type:text"`
}

// ReleaseUsersJobs - 用户任务关联
type ReleaseUserJob struct {
	ID     int64 `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	JobId  int64 `gorm:"column:job_id;type:bigint(20);not null"`
	UserId int64 `gorm:"column:user_id;type:bigint(20);not null"`
}

type ReleasePipeline struct {
	ID           int64     `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	JobId        int64     `gorm:"column:job_id;type:bigint(20);not null"`
	BuildId      int64     `gorm:"column:build_id;type:bigint(20);not null"`
	StartDate    time.Time `gorm:"column:start_date;type:datetime(6)"`
	EndDate      time.Time `gorm:"column:end_date;type:datetime(6)"`
	DurationDate time.Time `gorm:"column:duration_date;type:datetime(6)"`
	Status       bool      `gorm:"column:status;type:tinyint(1);not null;default:0"`
	UserID       int64     `gorm:"column:user_id;type:bigint(20);not null"`
}

// OmsJenkinsInstances - jenkins实例
type OmsJenkinsInstances struct {
	ID       int64  `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	EnvID    int64  `gorm:"column:env_id;type:bigint(20);not null"`
	Name     string `gorm:"column:name;type:varchar(150);unique;not null"`
	Address  string `gorm:"column:address;type:varchar(150);not null"`
	AuthType bool   `gorm:"column:auth_type;type:tinyint(1);not null"`
	UserName string `gorm:"column:username;type:varchar(150);not null"`
	Password string `gorm:"column:password;type:varchar(150);not null"`
	Desc     string `gorm:"column:desc;type:text"`
}
