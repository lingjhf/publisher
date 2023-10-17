package utils

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// 连接sqlite数据库
func ConnectSqlite(dsn string, config *gorm.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn), config)
	if err != nil {
		panic("fail to connect database")
	}
	return db
}

func HostString(host string, port uint) string {
	return fmt.Sprintf("%s:%d", host, port)
}
