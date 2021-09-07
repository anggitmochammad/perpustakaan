package models

// variabel harus besar agar dapat diimport
type Buku struct{
	// untuk dipanggil| tipe data | untuk database
	ID int  `gorm:"PRIMARY_KEY"`
	NamaBuku string 	
	NoTelp string 	
	Stok int 
}

type Member struct{
	ID int  `gorm:"PRIMARY_KEY"`
	Nama string 	
	Jenis_kelamin string 
	// has many
	Transaksi []Transaksi `gorm:"foreignkey:ID_Member"`
}

type Transaksi struct{
	ID int  `gorm:"PRIMARY_KEY"`
	ID_Member uint
	Tipe string 
	// has many
	DetailTransaksi []DetailTransaksi `gorm:"foreignkey:ID_Transaksi"`
}

type DetailTransaksi struct{
	ID int  `gorm:"PRIMARY_KEY"`
	ID_Transaksi uint
	ID_Buku uint
	Jumlah int
}