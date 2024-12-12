package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/oms/settings"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
	// 连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}

var gdb *gorm.DB

// gorm
func InitDB(cfg *settings.MySQLConfig) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("failed to connect mysql database", zap.Error(err))
		return
	}
	zap.L().Info("success to connect mysql database")
	// sqlDB, _ := db.DB()
	// sqlDB.SetMaxIdleConns(cfg.MaxOpenConns)
	// sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	gdb = db
}

func CloseDB() {
	sqlDB, err := gdb.DB()
	if err != nil {
		zap.L().Error("failed to get database instance", zap.Error(err))
		return
	}
	if err := sqlDB.Close(); err != nil {
		zap.L().Error("failed to close database", zap.Error(err))
	}
}
