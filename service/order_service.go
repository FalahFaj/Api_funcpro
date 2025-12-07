package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"projek_funcpro_kel12/model"
	"projek_funcpro_kel12/repository"
	"time"
)

type CreateOrderInput struct {
	Items []OrderItemInput `json:"items" binding:"required"`
	// TotalHarga float64 `json:"total_harga" binding:"required"`
	// Status string `json:"status" binding:"required"`
}

type OrderItemInput struct {
	ProdukID int64 `json:"produk_id" binding:"required"`
	Jumlah   int   `json:"jumlah" binding:"required,min=1"`
}

type OrderService interface {
	CreateOrder(ctx context.Context, pembeliId int64, input CreateOrderInput) (*model.Order, error)
	GetAllOrder(ctx context.Context) ([]model.Order, error)
	GetOrderById(ctx context.Context, id int64) (*model.Order, error)
}

type orderService struct {
	orderRepo  repository.OrderRepository
	produkRepo repository.ProdukRepository
	db         *sql.DB
}

func NewOrderService(orderRepo repository.OrderRepository, produkRepo repository.ProdukRepository, db *sql.DB) *orderService {
	return &orderService{orderRepo, produkRepo, db}
}

func (s *orderService) CreateOrder(ctx context.Context, pembeliId int64, input CreateOrderInput) (*model.Order, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var totalHarga int64
	var items []model.OrderItem

	for _, itemInput := range input.Items {
		produk, err := s.produkRepo.GetProdukByIdForUpdate(ctx, tx, itemInput.ProdukID)
		if err != nil {
			return nil, errors.New("produk tidak ditemukan")
		}
		if produk.Stok < int(itemInput.Jumlah) {
			return nil, fmt.Errorf("stok produk %s kurang, sisa %d", produk.NamaProduk, produk.Stok)
		}

		item := model.OrderItem{
			OrderId:           0,
			ProdukId:          itemInput.ProdukID,
			Jumlah:            int64(itemInput.Jumlah),
			HargaKetikaDIBeli: float64(produk.Harga), // Tetap float64 jika diperlukan untuk presisi
		}
		items = append(items, item)
		totalHarga += produk.Harga * int64(itemInput.Jumlah)

		produk.Stok -= int(itemInput.Jumlah)
		err = s.produkRepo.UpdateStok(ctx, tx, produk)
		if err != nil {
			return nil, err
		}
	}

	order := &model.Order{
		PembeliId:  pembeliId,
		TotalHarga: totalHarga,
		Status:     "pending",
		CreatedAt:  time.Now(),
		Items:      items,
	}

	orderId, err := s.orderRepo.CreateHeader(ctx, tx, order)
	if err != nil {
		return nil, err
	}

	order.Id = orderId

	for _, item := range items {
		item.OrderId = orderId
		err = s.orderRepo.CreateItem(ctx, tx, &item)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderService) GetAllOrder(ctx context.Context) ([]model.Order, error) {
	orders, err := s.orderRepo.GetAllOrder(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *orderService) GetOrderById(ctx context.Context, id int64) (*model.Order, error) {
	order, err := s.orderRepo.GetOrderById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order tidak ditemukan")
		}
		return nil, err
	}
	return order, nil
}
