package databases

import (
	"fmt"
	"relasi-gorm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const(
	username = "root"
	password = "123456"
	database = "databasego"
	hostname = "localhost"
	portname = "3306"
)

var DB *gorm.DB

func DatabaseInit() error {
	var err error
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",username,password,hostname,portname,database)
	// fmt.Println(dsn)
	// DB, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/databasego?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
        PrepareStmt:            true,
	})

	if err != nil {
		fmt.Println("hahahaarifin")
		panic("failed to connect database")
	}

	if err != nil {
		fmt.Println("hahaha")
		panic(err)
	}

	DB.AutoMigrate(&models.User{},&models.Locker{}, &models.Post{})
	return nil	
}