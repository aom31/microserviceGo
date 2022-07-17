package main

import (
	"net/http"

	"github.com/aom31/restapicrud/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root@tcp(127.0.0.1:3306)/goapilog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}

	r := gin.Default()
	r.GET("/customers", func(c *gin.Context) {
		var customers []model.Customer
		db.Find(&customers)
		c.JSON(http.StatusOK, customers)
	})

	r.GET("/customers/:id", func(c *gin.Context) {
		id := c.Param("id")

		var customer model.Customer
		db.First(&customer, id)
		c.JSON(http.StatusOK, customer)
	})
	r.POST("/customers", func(c *gin.Context) {
		//binding json
		var jsonCustomer model.Customer
		if err := c.ShouldBindJSON(&jsonCustomer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := db.Create(&jsonCustomer)
		c.JSON(200, gin.H{"RowAffected": result.RowsAffected})
	})

	r.DELETE("/customers/:id", func(c *gin.Context) {
		//find old data in db
		id := c.Param("id")
		var customer model.Customer
		db.First(&customer, id)

		db.Delete(&customer)
		c.JSON(200, customer)
	})

	r.PUT("/customers/:id", func(c *gin.Context) {
		var customer model.Customer
		var updateCustomer model.Customer
		if err := c.ShouldBindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.First(&updateCustomer, customer.ID)
		updateCustomer.Firstname = customer.Firstname
		updateCustomer.Lastname = customer.Lastname
		updateCustomer.Email = customer.Email
		updateCustomer.Telnumber = customer.Telnumber
		updateCustomer.IDcard = customer.IDcard

		db.Save(updateCustomer)
		c.JSON(200, updateCustomer)
	})
	r.Run()

}
