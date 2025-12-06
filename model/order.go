package model

import "time"

type Order struct {
	Id         int64       `json:"id" db:"id"`
	PembeliId  int64       `json:"pembeli_id" db:"pembeli_id"`
	TotalHarga int64       `json:"total_harga" db:"total_harga"`
	Status     string      `json:"status" db:"status"`
	CreatedAt  time.Time   `json:"created_at" db:"created_at"`
	Barang2    []OrderItem `json:"barang2" db:"-"`
}
