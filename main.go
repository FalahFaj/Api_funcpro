package main

import (
	"log"

	"projek_funcpro_kel12/config"
)

func main() {
	dsn, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Gagal memuat konfigurasi: %v", err)
	}
	log.Println("Konfigurasi berhasil dimuat.")

	db, err := config.NewConnection(dsn)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}
	defer db.Close()
	log.Println("Koneksi database berhasil.")

	log.Println("Server berjalan di port :8080")
}
