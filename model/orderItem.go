package model

type OrderItem struct {
	Id        int64 `json:"id" db:"id"`
	OrderId   int64 `json:"order_id" db:"order_id"`
	ProdukId  int64 `json:"produk_id" db:"produk_id"`
	Jumlah    int64 `json:"jumlah" db:"jumlah"`
	HargaKetikaDIBeli float64 `json:"harga_ketika_dibeli" db:"harga_ketika_dibeli"`
}