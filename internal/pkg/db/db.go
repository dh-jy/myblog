package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// MySQLOptions 定义 MySQL 数据库的选项.
type MySQLOptions struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
}

// DSN 从 MySQLOptions 返回 DSN.
func (o *MySQLOptions) DSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		o.Username,
		o.Password,
		o.Host,
		o.Database,
		true,
		"Local")
}

// NewMySQL 使用给定的选项创建一个新的 gorm 数据库实例.
func NewMySQL(opts *MySQLOptions) (*gorm.DB, error) {
	//logLevel := logger.Silent
	//if opts.LogLevel != 0 {
	//	logLevel = logger.LogLevel(opts.LogLevel)
	//}
	db, err := gorm.Open("mysql", opts.DSN())
	//db, err := gorm.Open(mysql.Open(opts.DSN()), &gorm.Config{
	//	Logger: logger.Default.LogMode(logLevel),
	//})
	if err != nil {
		return nil, err
	}
	//
	sqlDB := db.DB()

	//if err != nil {
	//	return nil, err
	//}

	// SetMaxOpenConns 设置到数据库的最大打开连接数
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

	// SetConnMaxLifetime 设置连接可重用的最长时间
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	// SetMaxIdleConns 设置空闲连接池的最大连接数
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	//db.AutoMigrate(&model.UserM{})
	//db.AutoMigrate(&model.PostM{})

	return db, nil
}
