package main

import (
	"bioskop-app-adi/database"
	"bioskop-app-adi/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	database.ConnectDB()

	// Inisialisasi router Gin
	router := gin.Default()

	// Registrasi route
	routers.RegisterBioskopRoutes(router)

	// Jalankan web server
	router.Run(":8080")
}
