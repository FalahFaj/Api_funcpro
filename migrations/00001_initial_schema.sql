-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE produk (
    id SERIAL PRIMARY KEY,
    petani_id INT NOT NULL,
    nama_produk VARCHAR(255) NOT NULL,
    deskripsi TEXT,
    harga DECIMAL(12, 2) NOT NULL CHECK (harga >= 0),
    stok INT NOT NULL DEFAULT 0 CHECK (stok >= 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT fk_petani
        FOREIGN KEY(petani_id) 
        REFERENCES users(id)
        ON DELETE CASCADE 
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    pembeli_id INT NOT NULL,
    total_harga DECIMAL(12, 2) NOT NULL CHECK (total_harga >= 0),
    status VARCHAR(50) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'completed', 'cancelled')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT fk_pembeli
        FOREIGN KEY(pembeli_id) 
        REFERENCES users(id)
        ON DELETE RESTRICT 
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    produk_id INT NOT NULL,
    jumlah INT NOT NULL CHECK (jumlah > 0),
    harga_ketika_dibeli DECIMAL(12, 2) NOT NULL, 
    CONSTRAINT fk_order
        FOREIGN KEY(order_id) 
        REFERENCES orders(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_produk
        FOREIGN KEY(produk_id) 
        REFERENCES produk(id)
        ON DELETE RESTRICT 
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS produk;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
