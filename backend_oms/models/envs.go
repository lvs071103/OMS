package models

// Env - 环境
type OmsEnvConfig struct {
	ID    int64  `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	Name  string `gorm:"column:name;type:varchar(100);unique;not null"`
	Label string `gorm:"column:label;type:varchar(100);not null"`
	Desc  string `gorm:"column:desc;type:text"`
}
