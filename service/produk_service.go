package service

import (
	"context"
	"projek_funcpro_kel12/model"
	"projek_funcpro_kel12/repository"
	"time"
)

type InputProduk struct {
	NamaProduk string  `json:"nama_produk" binding:"required"`
	Deskripsi  string  `json:"deskripsi" binding:"required"`
	Harga      int64 `json:"harga" binding:"required"`
	Stok       int     `json:"stok" binding:"required"`
}

type ProdukService interface {
	CreateProduk(ctx context.Context, petaniId int64, input InputProduk) (*model.Produk, error)
	GetAllProduk(ctx context.Context) ([]model.Produk, error)
	GetProdukById(ctx context.Context, id int64) (*model.Produk, error)
	UpdateProduk(ctx context.Context, id int64, input InputProduk) (*model.Produk, error)
	DeleteProduk(ctx context.Context, id int64) error
}

type produkService struct {
	produkRepo repository.ProdukRepository
}

func NewProdukService(produkRepo repository.ProdukRepository) *produkService {
	return &produkService{produkRepo}
}

func (s *produkService) CreateProduk(ctx context.Context, petaniId int64, input InputProduk) (*model.Produk, error) {
	produk := &model.Produk{
		PetaniId:   petaniId,
		NamaProduk: input.NamaProduk,
		Deskripsi:  input.Deskripsi,
		Harga:      input.Harga,
		Stok:       input.Stok,
		CreatedAt:  time.Now(),
	}

	id, err := s.produkRepo.Tambah(ctx, produk)
	if err != nil {
		return nil, err
	}

	produk.Id = id
	return produk, nil
}

func (s *produkService) GetAllProduk(ctx context.Context) ([]model.Produk, error) {
	produks, err := s.produkRepo.GetAllProduk(ctx)
	if err != nil {
		return nil, err
	}
	return produks, nil
}

func (s *produkService) DeleteProduk(ctx context.Context, id int64) error {
	err := s.produkRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *produkService) GetProdukById(ctx context.Context, id int64) (*model.Produk, error) {
	produk, err := s.produkRepo.GetProdukById(ctx, id)
	if err != nil {
		return nil, err
	}
	return produk, nil
}

func (s *produkService) UpdateProduk(ctx context.Context, id int64, input InputProduk) (*model.Produk, error) {
	produk, err := s.produkRepo.GetProdukById(ctx, id)
	if err != nil {
		return nil, err
	}

	produk.NamaProduk = input.NamaProduk
	produk.Deskripsi = input.Deskripsi
	produk.Harga = input.Harga
	produk.Stok = input.Stok

	err = s.produkRepo.Update(ctx, produk)
	if err != nil {
		return nil, err
	}

	return produk, nil
}
