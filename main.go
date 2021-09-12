package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"perpustakaan/controllers"
	"perpustakaan/middleware"
	"perpustakaan/models"
)

func main()  {
	// inisialisasi route
	route := gin.Default()

	// model
	db := models.SetupModels()
	route.Use(func(c *gin.Context) {
		c.Set("db",db)
		// jalankan proses berikutnya
		c.Next()
	})

	// membuat route
	route.GET("/",func(c *gin.Context) {
		// mengirim status dan data json
		c.JSON(http.StatusOK, gin.H{"data" :"Test Api" })
	})

	route.POST("/registrasi",controllers.CreateUser)
	route.POST("/login",controllers.Login)

	buku := route.Group("buku",middleware.Auth)
	{
		buku.GET("/",controllers.GetAllBuku)
		buku.POST("/",controllers.CreateBuku)
		buku.PUT("/:id",controllers.UpdateBuku)
		buku.DELETE("/:id",controllers.DeleteBuku)
	}
	member := route.Group("member")
	{
		member.GET("/",controllers.GetAllMember)
		member.POST("/",controllers.CreateMember)
		member.PUT("/:id",controllers.UpdateMember)
		member.DELETE("/:id",controllers.DeleteMember)
	}
	transaksi := route.Group("transaksi")
	{
		transaksi.GET("/",controllers.GetAllTransaksi)
		transaksi.POST("/",controllers.CreateTransaksi)
		transaksi.PUT("/:id",controllers.UpdateTransaksi)
		transaksi.DELETE("/:id",controllers.DeleteTransaksi)
	}

	// run route . Jika tidak diisi defaultnya localhost:8080
	route.Run(":7070")
}
