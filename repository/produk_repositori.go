package repository

import (
	"context"
	"database/sql"
	"projek_funcpro_kel12/model"
)

type ProdukRepository interface {
	Tambah(ctx context.Context, produk *model.Produk) (int64, error)
	GetAllProduk(ctx context.Context) (*model.Produk)
	GetProdukById(ctx context.Context, id int64) (*model.Produk, error)
	GetProdukByName(ctx context.Context, name string) (*model.Produk, error)
}

type produkRepository struct {
	db *sql.DB
}

func NewProdukRepository(db *sql.DB) *produkRepository {
	return &produkRepository{db}
}

func (r *produkRepository) GetAllProduk(ctx context.Context) (*model.Produk, error) {
	var produks model.Produk
	query := `SELECT nama_produk, deskripsi, harga, stok, created_at FROM produk`
	err := r.db.QueryRow(query).Scan(&produks.NamaProduk, &produks.Deskripsi, &produks.Harga, &produks.Stok, &produks.CreatedAt)

	if err != nil{
		return nil, err
	}
	return &produks, nil
}

