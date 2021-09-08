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
	ID_Transaksi uint
	ID_Buku uint
	Jumlah int
}

// type detailtransaksicreate struct{
// 	ID_Transaksi uint
// 	ID_Buku uint
// 	Jumlah int
// }

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

	var member models.Member

	// input validasi
	if err := c.ShouldBindJSON(&dataInput);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
	}


	// proses input
	if dataInput.ID_Member == 0 || dataInput.Tipe == "" {
		c.JSON(http.StatusBadRequest,gin.H{"error" : "data error"})
		return
	}

	

	//  memasukkan input ke struct
	transaksi := models.Transaksi{
		// Model : data input
		ID_Member: dataInput.ID_Member,
		Tipe: dataInput.Tipe,
	}

		// cek foreign key member
	//// SELECT * FROM members ORDER BY id LIMIT 1;
	if err := db.Where("id = ?", dataInput.ID_Member).First(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"anggota tidak ditemukan"})
		return
	}

	// membuat transaksi
	db.Create(&transaksi)

	// deklarasi model buku
	var buku models.Buku

	detailtransaksi := models.DetailTransaksi{
		ID_Transaksi: dataInput.ID_Transaksi,
		ID_Buku: dataInput.ID_Buku,
		Jumlah: dataInput.Jumlah,
	}

	// cek id buku
	if err := db.Where("id = ?", dataInput.ID_Buku).First(&buku).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "buku tidak ditemukan"})
			return
	}
	
	// cek stok
	if dataInput.Tipe == "peminjaman" && buku.Stok < dataInput.Jumlah {
		c.JSON(http.StatusBadRequest, gin.H{"error":"permintaan lebih banyak dari stok"})
		return
	}

	// membuat detail transaksi
	db.Create(&detailtransaksi)

	// update stok
	if dataInput.Tipe == "peminjaman" {
		buku.Stok = buku.Stok - dataInput.Jumlah
		db.Model(&buku).Update(buku)
	}
	if dataInput.Tipe == "pengembalian" {
		buku.Stok = buku.Stok + dataInput.Jumlah
		db.Model(&buku).Update(buku)
	}
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
