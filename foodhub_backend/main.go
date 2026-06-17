package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//* Global variable untuk koneksi ke database
var DB *gorm.DB

//* Membuat custom Middleware logger (Gin)
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context)  {
		t := time.Now()

		//* Proses request (melanjutkan ke handler utama)
		c.Next()

		//* Setelah handler selesai hitung waktu eksekusi
		latency := time.Since(t)
		status := c.Writer.Status()

		log.Printf("[FoodHub-log] %s %s %s %s", c.Request.Method, status, c.Request.URL.Path, latency)
	}
}


//* Fungsi Untuk Koneksi ke Database
func initDB() {
	//* sesuai dengan data yang sudah di buat di postgresql
	dsn := "host=localhost user=foodhub_user password=foodhub123 dbname=foodhub_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	var err error

	//* Membukan koneksi menggunakan GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database")
	}

	fmt.Println("Berhasil terkoneksi ke database PostgreSQL")
}

func main() {
	//* Inisialisasi Databse
	initDB()

	//* Inisialisasi Router Gin
	r := gin.New()

	//* Memasang middleware Global
	r.Use(GinLogger())
	r.Use(gin.Recovery()) //* Mencegah server crash jika terjadi panic

	//* Route utama (endpoint API)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"message": "Selamat datang di API foodhub marketplace dengan Gin Gonic!",
		})
	})

	//* Menjalankan server di port :8080
	r.Run(":8080")
}