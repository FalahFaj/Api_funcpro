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

	// 3. Inisialisasi Layer (Wiring) - Contoh untuk User
	// userRepo := repository.NewUserRepository(db)
	// userService := service.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userService)

	// 4. Setup Router & Jalankan Server
	// router := http.NewServeMux()
	// router.HandleFunc("/users", userHandler.CreateUser)

	log.Println("Server berjalan di port :8080")
	// if err := http.ListenAndServe(":8080", router); err != nil {
	// 	log.Fatalf("Gagal menjalankan server: %v", err)
	// }
}
