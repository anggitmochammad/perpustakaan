package controllers

import (
	"net/http"

	"perpustakaan/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type membercreate struct{
	Nama string `json:"nama"`
	Jenis_kelamin string `json:"jenis_kelamin"`
}

// nama harus besar agar dapat dipanggil
func GetAllMember(c *gin.Context){
	// koneksi db . membuka sesuai setingan setup.go
	db:= c.MustGet("db").(*gorm.DB)

	var member []models.Buku
	// query
	db.Find(&member)
	c.JSON(http.StatusOK,gin.H{"data" : member})
} 

func CreateMember(c *gin.Context)  {
	// koneksi db . membuka sesuai setingan setup.go
	db:= c.MustGet("db").(*gorm.DB)

	// input validasi
	var dataInput membercreate
	if err := c.ShouldBindJSON(&dataInput);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
	}

	// proses input
	if dataInput.Nama == "" || dataInput.Jenis_kelamin == ""{
		c.JSON(http.StatusBadRequest,gin.H{"error" : "data tidak boleh kosong"})
		return
	}
	member := models.Member{
		// Model : data input
		Nama: dataInput.Nama,
		Jenis_kelamin: dataInput.Jenis_kelamin,
	}

	db.Create(&member)
	c.JSON(http.StatusOK,gin.H{"status" : "Berhasil Ditambahkan"})
	
}

func UpdateMember(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	// cek data sesuai id
	var member models.Member
	if err := db.Where("id = ?", c.Param("id")).First(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"data tidak ditemukan"})
		return
	}

	// input validasi
	var dataInput membercreate
	if err := c.ShouldBindJSON(&dataInput);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
	}

	// proses input
	if dataInput.Nama == "" || dataInput.Jenis_kelamin == "" {
		c.JSON(http.StatusBadRequest,gin.H{"error" : "data tidak boleh kosong"})
		return
	}
	db.Model(&member).Update(dataInput)
	
	c.JSON(http.StatusOK,gin.H{"status" : "Berhasil Diubah"})

}

func DeleteMember(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)
	
	var buku models.Transaksi
	if err := db.Where("id = ?", c.Param("id")).First(&buku).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"data tidak ditemukan"})
		return
	}

	db.Delete(&buku)
	c.JSON(http.StatusOK,gin.H{"status" : "Berhasil Dihapus"})
}
