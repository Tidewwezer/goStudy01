package common

import (
	"fmt"
	"ginessential/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "42.192.40.207"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "Wwezer111555"
	charset := "utf8"
	fmt.Println(host)
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})

	return db
}

func GetDB() *gorm.DB {
	DB = InitDB()
	return DB
}
