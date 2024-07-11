package repositories

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var m sync.Once

func InitDatabase() {
	m.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "root2024", "127.0.0.1", "3306", "admin")
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 禁用表名复数
			},
		})
		DB = db
		if err != nil {
			log.Fatal("failed to connect database", err)
		}
	})
}
