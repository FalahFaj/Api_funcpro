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

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, jwtSecret)
	userHandler := handler.NewUserHandler(userService)

	produkRepo := repository.NewProdukRepository(db)
	produkService := service.NewProdukService(produkRepo)
	produkHandler := handler.NewProdukHandler(produkService)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, produkRepo, db)
	orderHandler := handler.NewOrderHandler(orderService)

	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)

	http.HandleFunc("GET /produk", produkHandler.GetAllProduk)

	produkCreateHandler := http.HandlerFunc(produkHandler.CreateProduk)
	protectedProdukCreate := handler.AuthMiddleware(handler.RoleMiddleware(produkCreateHandler, "petani"), userService, jwtSecret)
	http.Handle("POST /produk", protectedProdukCreate)

	produkByIdHandler := http.HandlerFunc(produkHandler.KelolaProdukById)
	protectedProdukById := handler.AuthMiddleware(handler.RoleMiddleware(produkByIdHandler, "petani", "pembeli"), userService, jwtSecret)
	http.Handle("/produk/{id}", protectedProdukById)

	userManagementHandler := http.HandlerFunc(userHandler.KelolaAkun)
	protectedUserManagement := handler.AuthMiddleware(userManagementHandler, userService, jwtSecret)
	http.Handle("/users/{id}", protectedUserManagement)

	orderCreateHandler := http.HandlerFunc(orderHandler.CreateOrder)
	protectedOrderCreate := handler.AuthMiddleware(handler.RoleMiddleware(orderCreateHandler, "pembeli"), userService, jwtSecret)
	http.Handle("POST /orders", protectedOrderCreate)

	orderGetHandler := http.HandlerFunc(orderHandler.GetOrderById)
	protectedOrderGet := handler.AuthMiddleware(orderGetHandler, userService, jwtSecret)
	http.Handle("GET /orders/{id}", protectedOrderGet)

	port := "8080"
	addr := ":" + port
	log.Printf("Server berjalan di http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
