package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb(username string, password string, host string, port string, dbname string, minConn int, maxConn int) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dbConf, _ := db.DB()
	dbConf.SetMaxIdleConns(minConn)
	dbConf.SetMaxOpenConns(maxConn)
	DB = db
}
