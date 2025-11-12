package main

import (
	"log"
	"net/http"

	"projek_funcpro_kel12/config"
	"projek_funcpro_kel12/handler"
	"projek_funcpro_kel12/repository"
	"projek_funcpro_kel12/service"
)

func main() {
	dsn, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Gagal memuat konfigurasi: %v", err)
	}
	db, err := config.NewConnection(dsn)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}
	defer db.Close()

	jwtSecret, err := config.LoadJWTSecret()
	if err != nil {
		log.Fatalf("Gagal memuat JWT secret: %v", err)
	}

	// Inisialisasi semua layer
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, jwtSecret)
	userHandler := handler.NewUserHandler(userService)

	// Mendaftarkan rute untuk user
	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)

	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
