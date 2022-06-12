package dao

import (
	"action-game/cmd/api/internal/config"
	"action-game/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

func InitDB() {
	MysqlCfg := config.GetUserCfg().MysqlCfg
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       MysqlCfg.DataSource, // DSN data source name
		DefaultStringSize:         171,                 // string 类型字段的默认长度
		SkipInitializeWithVersion: false,               // 根据当前 MySQL 版本自动配置, &gorm.Config{})
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 表名为复数
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 逻辑外键
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDB, err := db.DB()

	// SetMaxIdleCons 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
	DB.AutoMigrate(&model.Room{})
	// db.First(&model.UserInfo{})
	fmt.Println("数据库链接成功")
}
