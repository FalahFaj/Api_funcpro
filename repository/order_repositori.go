package repository

import (
	"context"
	"database/sql"
	"projek_funcpro_kel12/model"
)

type OrderRepository interface {
	Tambah(ctx context.Context, order *model.Order) (int64, error)
	GetAllOrder(ctx context.Context) ([]model.Order, error)
	GetOrderById(ctx context.Context, id int64) (*model.Order, error)
	Update(ctx context.Context, order *model.Order) error
	Delete(ctx context.Context, id int64) error
	CreateHeader(ctx context.Context, tx *sql.Tx, order *model.Order) (int64, error)
	CreateItem(ctx context.Context, tx *sql.Tx, item *model.OrderItem) (int64, error)
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Tambah(ctx context.Context, order *model.Order) (int64, error) {
	var id int64
	query := `INSERT INTO orders pembeli_id, total_harga, status, created_at VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.db.QueryRow(query, order.PembeliId, order.TotalHarga, order.Status, order.CreatedAt).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *orderRepository) GetAllOrder(ctx context.Context) ([]model.Order, error) {
	var orders []model.Order
	query := `SELECT id, pembeli_id, total_harga, status, created_at FROM orders`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.Id, &order.PembeliId, &order.TotalHarga, &order.Status, &order.CreatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepository) GetOrderById(ctx context.Context, id int64) (*model.Order, error) {
	var order model.Order
	query := `SELECT id, pembeli_id, total_harga, status, created_at FROM orders WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&order.Id, &order.PembeliId, &order.TotalHarga, &order.Status, &order.CreatedAt)
	if err != nil {
		return nil, err
	}

	queryItems := `SELECT id, order_id, produk_id, jumlah, harga_ketika_dibeli FROM order_items WHERE order_id = $1`
	rows, err := r.db.QueryContext(ctx, queryItems, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.OrderItem
		if err := rows.Scan(&item.Id, &item.OrderId, &item.ProdukId, &item.Jumlah, &item.HargaKetikaDIBeli); err != nil {
			return nil, err
		}
		order.Items = append(order.Items, item)
	}

	return &order, nil
}

func (r *orderRepository) Update(ctx context.Context, order *model.Order) error {
	query := `UPDATE orders SET pembeli_id = $1, total_harga = $2, status = $3, created_at = $4 WHERE id = $5`
	hasil, err := r.db.ExecContext(ctx, query, order.PembeliId, order.TotalHarga, order.Status, order.CreatedAt, order.Id)
	if err != nil {
		return err
	}
	if rowsAffected, _ := hasil.RowsAffected(); rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *orderRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM orders WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *orderRepository) CreateHeader(ctx context.Context, tx *sql.Tx, order *model.Order) (int64, error) {
	var id int64
	query := `INSERT INTO orders (pembeli_id, total_harga, status, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := tx.QueryRowContext(ctx, query, order.PembeliId, order.TotalHarga, order.Status, order.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *orderRepository) CreateItem(ctx context.Context, tx *sql.Tx, item *model.OrderItem) (int64, error) {
	query := `INSERT INTO order_items (order_id, produk_id, jumlah, harga_ketika_dibeli) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int64
	err := tx.QueryRowContext(ctx, query, item.OrderId, item.ProdukId, item.Jumlah, item.HargaKetikaDIBeli).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
