package controllers

import (
	"net/http"

	"perpustakaan/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type transaksicreate struct{
	ID_Member uint
	Tipe string 
}

type detailtransaksicreate struct{
	ID_Transaksi uint
	ID_Buku uint
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

	// deklarasi struct untuk input
	var dataInput transaksicreate
	var detailTransaksiInput detailtransaksicreate

	// deklarasi model buku
	var member models.Member
	var buku models.Buku
	var detailtransaksi models.DetailTransaksi

	// input validasi
	if err := c.ShouldBindJSON(&dataInput);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
	}


	// proses input
	if dataInput.ID_Member == 0 || dataInput.Tipe == "" {
		c.JSON(http.StatusBadRequest,gin.H{"error" : "data error"})
		return
	}

	// cek foreign key member
	//// SELECT * FROM members ORDER BY id LIMIT 1;
	if err := db.Where("id = ?", dataInput.ID_Member).First(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"anggota tidak ditemukan"})
		return
	}

	// cek id buku
	if err := db.Where("id = ?", detailTransaksiInput.ID_Buku).First(&buku).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"buku tidak ditemukan"})
			return
		}
	
	// cek stok
	if dataInput.Tipe == "pengembalian" && buku.Stok < detailTransaksiInput.Jumlah {
		c.JSON(http.StatusBadRequest, gin.H{"error":"permintaan lebih banyak dari stok"})
		return
	}
	
	transaksi := models.Transaksi{
		// Model : data input
		ID_Member: dataInput.ID_Member,
		Tipe: dataInput.Tipe,
	}

	// membuat transaksi
	db.Create(&transaksi)
	// membuat detail transaksi
	db.Model(&detailtransaksi).Update(detailTransaksiInput)
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
	if dataInput.ID_Member == 0 || dataInput.Tipe == "" {
		c.JSON(http.StatusBadRequest,gin.H{"error" : "data tidak boleh kosong"})
		return
	}

	// update database
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
