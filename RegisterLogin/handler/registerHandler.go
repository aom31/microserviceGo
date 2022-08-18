package handler

import (
	"net/http"

	"github.com/aom31/regislog/constant"
	"github.com/aom31/regislog/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Register struct {
	Username string `json:"username" binding: "required"`
	Password string `json : "password" binding : "required"`
	Fullname string `json: "fullname" binding : "required"`
	Email    string `json: "email" binding: "required"`
}

func RegisterUser(c *gin.Context) {
	database.ConnectDb()

	var jsonRegis Register
	if err := c.ShouldBindJSON(&jsonRegis); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//check user create id yet?
	var userExist database.User
	database.Db.Where("username = ?", jsonRegis.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusBadGateway, gin.H{"status": "Error", "message": "User Exists!!"})
		return
	}
	//create users
	encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(jsonRegis.Password), constant.DefaultCost)
	user := database.User{Username: jsonRegis.Username, Password: string(encryptPassword), Fullname: jsonRegis.Fullname, Email: jsonRegis.Email}
	database.Db.Create(&user)
	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"Status": "OK", "message": "User created successful!", "userId": user.ID})
	} else {
		c.JSON(http.StatusOK, gin.H{"Status": "Fail", "message": "User that created already!!"})
	}

}
