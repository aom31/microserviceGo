package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//create struct of  db
type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Email    string
}

var Db *gorm.DB
var err error

func ConnectDb() {

	dsn := "root:3143@tcp(127.0.0.1:3306)/godatabase?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}
	//migrate the schema db
	Db.AutoMigrate(&User{})

}
