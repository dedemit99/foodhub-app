package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//* Fungsi Middleware
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//* Menjalankan handler pertama yang dituju
		next(w, r)

		//* Setelah handler selesai, cetak log-nya ke terminal
		log.Printf("[%s] %s %s took %s", r.RemoteAddr, r.Method, r.URL.Path, time.Since(start))
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selamat datang di API FoodHub Marketplace!")
}

func main() {
	//* Membungkus handleHome dengan loggerMiddleware
	http.HandleFunc("/", loggerMiddleware(handleHome))

	//* Menjalankan server di port 8080
	fmt.Print("Server bacekend berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}