package repositories

import (
	"fmt"
	"github.com/walterwong1001/admin/global"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	cfg := global.CONFIG.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名复数
		},
	})

	if err != nil {
		log.Fatal("Fatal error: failed to connect database", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
