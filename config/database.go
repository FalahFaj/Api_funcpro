package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka koneksi database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("gagal memverifikasi koneksi ke database: %w", err)
	}

	return db, nil
}
