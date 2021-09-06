package models

// digunakan untuk menambahkan tabel pada database

import (
	// untuk query dan database
	"github.com/jinzhu/gorm"

	// untuk menyambungkan mysql
	// underscore berfungsi jika import tidak digunakan tidak akan error
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB  {

	// membuka database 
	// mysql, username:password@(di mana databasenya. Jika online isi dengan ip)/nama database
  db, err := gorm.Open("mysql","root:@(localhost)/perpustakaan?charset=utf8&parseTime=True&loc=Local")

  if err != nil {
	  panic("Gagal koneksi database")	
  }

//   deklrasi struct
  db.AutoMigrate(&Buku{})
  db.AutoMigrate(&Member{})
  db.AutoMigrate(&Transaksi{})

//   defer db.Close()
  return db
}



