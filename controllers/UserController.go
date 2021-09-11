package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"net/http"

	"perpustakaan/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type usercreate struct{
	Username string 
	Password string
}

// nama harus besar agar dapat dipanggil
func GetAllUser(c *gin.Context){
	// koneksi db . membuka sesuai setingan setup.go
	db:= c.MustGet("db").(*gorm.DB)

	var user []models.User
	// query
	db.Find(&user)
	c.JSON(http.StatusOK,gin.H{"data" : user})
} 

func CreateUser(c *gin.Context)  {
	// koneksi db . membuka sesuai setingan setup.go
	db:= c.MustGet("db").(*gorm.DB)

	var usermodel []models.User

	// input validasi
	var dataInput usercreate
	if err := c.ShouldBindJSON(&dataInput);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
	}

	// proses input
	if dataInput.Username == "" || dataInput.Password == "" {
		c.JSON(http.StatusBadRequest,gin.H{"error" : "data tidak boleh kosong"})
		return
	}  		

	hash := sha1.New()
	hash.Write([]byte(dataInput.Password))
	password := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	user := models.User{
		// Model : data input
		Username: dataInput.Username,
		Password: password,
	}

	if err := db.Create(&user); err.Error != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error" : "username sudah digunakan"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"status" : "Berhasil Ditambahkan"})
	
}