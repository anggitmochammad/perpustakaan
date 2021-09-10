package controllers

import (
	"crypto/sha1"
	"net/http"

	"perpustakaan/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type usercreate struct{
	Username string `gorm:"unique"`
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

	// cek unique username
	if result := db.Where("username = ?", dataInput.Username).First(&usermodel); result.Value != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error" : "Username sudah digunakan"})
		return
	}


	hash := sha1.New()
	hash.Write([]byte(dataInput.Password))
	password := hash.Sum(nil)
	// dataInput.Password = password
	user := models.User{
		// Model : data input
		Username: dataInput.Username,
		Password: string(password),
	}

	db.Create(&user)
	c.JSON(http.StatusOK,gin.H{"status" : "Berhasil Ditambahkan"})
	
}