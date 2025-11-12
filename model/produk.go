package model

import "time"

type Produk struct {
	Id         int64     `json:"id" db:"id"`
	PetaniId   int64     `json:"petani_id" db:"petani_id"`
	NamaProduk string    `json:"nama_produk" db:"nama_produk"`
	Deskripsi  string    `json:"deskripsi" db:"deskripsi"`
	Harga      int64     `json:"harga" db:"harga"`
	Stok       int64     `json:"stok" db:"stok"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}