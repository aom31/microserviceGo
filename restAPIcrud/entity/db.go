package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/aom31/restapicrud/model"
)



func main() {
	//connect db xampp
	dsn := "root@tcp(127.0.0.1:3306)/goapilog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}

	//Migrate the schema or create table database data init
	db.AutoMigrate(&model.Customer{})

	//create data in field
	db.Create(&model.Customer{Firstname: "thamakorn", Lastname: "ketnoi", Email: "thamakorn@gmail.com", Telnumber: 829163122, IDcard: 1709225408342})
	db.Create(&model.Customer{Firstname: "nuti", Lastname: "supawat", Email: "natnat@gmail.com", Telnumber: 748852992, IDcard: 330425408322})
}
