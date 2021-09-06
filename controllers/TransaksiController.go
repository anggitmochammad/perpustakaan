package controllers

import (
	"net/http"

	"perpustakaan/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type transaksicreate struct{
	ID_Member uint
	Jenis_kelamin string 	
	Tipe string 
	Jumlah int
}

// nama harus besar agar dapat dipanggil
func GetAllTransaksi(c *gin.Context){
	// koneksi db . membuka sesuai setingan setup.go
	db:= c.MustGet("db").(*gorm.DB)

	// array model
	var transaksi []models.Transaksi
	// query
	db.Find(&transaksi)
	c.JSON(http.StatusOK,gin.H{"data" : transaksi})
} 

func CreateTransaksi(c *gin.Context)  {
	// koneksi db . membuka sesuai setingan setup.go
	db:= c.MustGet("db").(*gorm.DB)

	// input validasi
	var dataInput transaksicreate
	if err := c.ShouldBindJSON(&dataInput);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
	}

	// proses input
	if dataInput.ID_Member == 0 || dataInput.Tipe == "" || dataInput.Jenis_kelamin == "" || dataInput.Jumlah == 0 {
		c.JSON(http.StatusBadRequest,gin.H{"error" : "data tidak boleh kosong"})
		return
	}
	transaksi := models.Transaksi{
		// Model : data input
		ID_Member: dataInput.ID_Member,
		Tipe: dataInput.Tipe,
		Jenis_kelamin: dataInput.Jenis_kelamin,
		Jumlah: dataInput.Jumlah,
	}

	db.Create(&transaksi)
	c.JSON(http.StatusOK,gin.H{"status" : "Berhasil Ditambahkan"})
	
}

func UpdateTransaksi(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	// cek data sesuai id
	var transaksi models.Transaksi
	if err := db.Where("id = ?", c.Param("id")).First(&transaksi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"data tidak ditemukan"})
		return
	}

	// input validasi
	var dataInput transaksicreate
	if err := c.ShouldBindJSON(&dataInput);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
	}

	// proses input
	if dataInput.ID_Member == 0 || dataInput.Tipe == "" || dataInput.Jenis_kelamin == "" {
		c.JSON(http.StatusBadRequest,gin.H{"error" : "data tidak boleh kosong"})
		return
	}
	db.Model(&transaksi).Update(dataInput)
	
	c.JSON(http.StatusOK,gin.H{"status" : "Berhasil Diubah"})

}

func DeleteTransaksi(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)
	
	var transaksi models.Transaksi
	if err := db.Where("id = ?", c.Param("id")).First(&transaksi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"data tidak ditemukan"})
		return
	}

	db.Delete(&transaksi)
	c.JSON(http.StatusOK,gin.H{"status" : "Berhasil Dihapus"})
}
