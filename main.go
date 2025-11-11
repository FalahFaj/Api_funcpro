package main

import (
	"fmt"
	"log"
	"projek_funcpro_kel12/config"
)

func main() {
	dsn, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Gagal memuat konfigurasi: %v", err)
	}
	fmt.Println("Wayahe projekan bang")
	fmt.Println("Koneksi string berhasil dibuat:", dsn)
}
