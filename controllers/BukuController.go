package controllers

import (
	"net/http"

	"perpustakaan/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type bukucreate struct{
	NamaBuku string `json:"nama_buku"`
	Stok int 		`json:"stok"`
}

// nama harus besar agar dapat dipanggil
func GetAllBuku(c *gin.Context){
	// koneksi db . membuka sesuai setingan setup.go
	db:= c.MustGet("db").(*gorm.DB)

	var buku []models.Buku
	// query
	db.Find(&buku)
	c.JSON(http.StatusOK,gin.H{"data" : buku})
} 

func CreateBuku(c *gin.Context)  {
	// koneksi db . membuka sesuai setingan setup.go
	db:= c.MustGet("db").(*gorm.DB)

	// input validasi
	var dataInput bukucreate
	if err := c.ShouldBindJSON(&dataInput);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
	}

	// proses input
	if dataInput.NamaBuku == "" || dataInput.Stok == 0 {
		c.JSON(http.StatusBadRequest,gin.H{"error" : "data tidak boleh kosong"})
		return
	}
	buku := models.Buku{
		// Model : data input
		NamaBuku: dataInput.NamaBuku,
		Stok: dataInput.Stok,
	}

	db.Create(&buku)
	c.JSON(http.StatusOK,gin.H{"status" : "Berhasil Ditambahkan"})
	
}

func UpdateBuku(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	// cek data sesuai id
	var buku models.Buku
	if err := db.Where("id = ?", c.Param("id")).First(&buku).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"data tidak ditemukan"})
		return
	}

	// input validasi
	var dataInput bukucreate
	if err := c.ShouldBindJSON(&dataInput);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
	}

	// proses input
	if dataInput.NamaBuku == "" || dataInput.Stok == 0 {
		c.JSON(http.StatusBadRequest,gin.H{"error" : "data tidak boleh kosong"})
		return
	}
	db.Model(&buku).Update(dataInput)
	
	c.JSON(http.StatusOK,gin.H{"status" : "Berhasil Diubah"})

}

func DeleteBuku(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)
	
	var buku models.Buku
	if err := db.Where("id = ?", c.Param("id")).First(&buku).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"data tidak ditemukan"})
		return
	}

	db.Delete(&buku)
	c.JSON(http.StatusOK,gin.H{"status" : "Berhasil Dihapus"})
}
