package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"perpustakaan/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type userlogin struct{
	Username string 
	Password string
}


func Login(c *gin.Context) {


	var dataInput userlogin
	var user models.User

	db:= c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&dataInput);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
	}

	if err := db.Where("username = ?",dataInput.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"data tidak ditemukan"})
		return
	}


	// hash password
	hash := sha1.New()
	hash.Write([]byte(dataInput.Password))
	password := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	if dataInput.Username !=  user.Username || password != user.Password{ 
		c.JSON(http.StatusInternalServerError, gin.H{"error":"username atau password salah"})
		return
	}

	// membuat token
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

