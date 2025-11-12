package model

import "time"

type Order struct {
	Id         int64     `json:"id" db:"id"`
	PembeliId  int64     `json:"pembeli_id" db:"pembeli_id"`
	TotalHarga int64     `json:"total_harga" db:"total_harga"`
	Status     string    `json:"status" db:"status"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	Barang2 []BarangDiBeli `json:"barang2" db:"-"`
}

type BarangDiBeli struct {
	Id        int64 `json:"id" db:"id"`
	OrderId   int64 `json:"order_id" db:"order_id"`
	ProdukId  int64 `json:"produk_id" db:"produk_id"`
	Jumlah    int64 `json:"jumlah" db:"jumlah"`
	HargaKetikaDIBeli float64 `json:"harga_ketika_dibeli" db:"harga_ketika_dibeli"`
}