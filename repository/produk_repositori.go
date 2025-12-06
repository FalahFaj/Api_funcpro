package repository

import (
	"context"
	"database/sql"
	"projek_funcpro_kel12/model"
)

type ProdukRepository interface {
	Tambah(ctx context.Context, produk *model.Produk) (int64, error)
	GetAllProduk(ctx context.Context) ([]model.Produk, error)
	GetProdukById(ctx context.Context, id int64) (*model.Produk, error)
	GetProdukByName(ctx context.Context, name string) (*model.Produk, error)
	GetProdukByIdForUpdate(ctx context.Context, tx *sql.Tx, id int64) (*model.Produk, error)
	UpdateStok(ctx context.Context, tx *sql.Tx, produk *model.Produk) error
	Update(ctx context.Context, produk *model.Produk) error
	Delete(ctx context.Context, id int64) error
}

type produkRepository struct {
	db *sql.DB
}

func NewProdukRepository(db *sql.DB) *produkRepository {
	return &produkRepository{db}
}

func (r *produkRepository) Tambah(ctx context.Context, produk *model.Produk) (int64, error) {
	var id int64
	query := `INSERT INTO produk (petani_id, nama_produk, deskripsi, harga, stok, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, produk.PetaniId, produk.NamaProduk, produk.Deskripsi, produk.Harga, produk.Stok, produk.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *produkRepository) GetAllProduk(ctx context.Context) ([]model.Produk, error) {
	var produks []model.Produk
	query := `SELECT id, petani_id, nama_produk, deskripsi, harga, stok, created_at FROM produk`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var produk model.Produk
		err := rows.Scan(
			&produk.Id,
			&produk.PetaniId,
			&produk.NamaProduk,
			&produk.Deskripsi,
			&produk.Harga,
			&produk.Stok,
			&produk.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		produks = append(produks, produk)
	}

	return produks, nil
}

func (r *produkRepository) GetProdukById(ctx context.Context, id int64) (*model.Produk, error) {
	var produk model.Produk
	query := `SELECT id, petani_id, nama_produk, deskripsi, harga, stok, created_at FROM produk WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&produk.Id, &produk.PetaniId, &produk.NamaProduk, &produk.Deskripsi, &produk.Harga, &produk.Stok, &produk.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &produk, nil
}

func (r *produkRepository) GetProdukByName(ctx context.Context, name string) (*model.Produk, error) {
	var produk model.Produk
	query := `SELECT id, petani_id, nama_produk, deskripsi, harga, stok, created_at FROM produk WHERE nama_produk = $1`
	err := r.db.QueryRowContext(ctx, query, name).Scan(&produk.Id, &produk.PetaniId, &produk.NamaProduk, &produk.Deskripsi, &produk.Harga, &produk.Stok, &produk.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &produk, nil
}

func (r *produkRepository) GetProdukByIdForUpdate(ctx context.Context, tx *sql.Tx, id int64) (*model.Produk, error) {
	var produk model.Produk
	query := `SELECT id, petani_id, nama_produk, deskripsi, harga, stok, created_at FROM produk WHERE id = $1 FOR UPDATE`
	err := tx.QueryRowContext(ctx, query, id).Scan(
		&produk.Id,
		&produk.PetaniId,
		&produk.NamaProduk,
		&produk.Deskripsi,
		&produk.Harga,
		&produk.Stok,
		&produk.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &produk, nil
}

func (r *produkRepository) UpdateStok(ctx context.Context, tx *sql.Tx, produk *model.Produk) error {
	query := `UPDATE produk SET stok = $1 WHERE id = $2`
	_, err := tx.ExecContext(ctx, query, produk.Stok, produk.Id)
	return err
}

func (r *produkRepository) Update(ctx context.Context, produk *model.Produk) error {
	query := `UPDATE produk SET nama_produk = $1, deskripsi = $2, harga = $3, stok = $4 WHERE id = $5`
	hasil, err := r.db.ExecContext(ctx, query, produk.NamaProduk, produk.Deskripsi, produk.Harga, produk.Stok, produk.Id)
	if err != nil {
		return err
	}
	if rowsAffected, _ := hasil.RowsAffected(); rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *produkRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM produk WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
