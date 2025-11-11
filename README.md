# Proyek API E-Commerce Pertanian (FuncPro - Kelompok 12)

Ini adalah backend API untuk platform e-commerce yang menghubungkan petani dengan pembeli. Proyek ini dibangun menggunakan Go dengan arsitektur berlapis (layered architecture) untuk memastikan kode yang bersih, terstruktur, dan mudah dikelola.

## Fitur

-   Manajemen Pengguna (Petani & Pembeli)
-   Manajemen Produk oleh Petani
-   Sistem Pemesanan oleh Pembeli
-   Riwayat Transaksi

## Arsitektur

Proyek ini mengadopsi arsitektur berlapis untuk memisahkan tanggung jawab:

-   **Handler/Controller**: Menerima request HTTP dan mengirimkan response.
-   **Service**: Berisi semua logika bisnis aplikasi.
-   **Repository**: Bertanggung jawab untuk berinteraksi dengan database.
-   **Model**: Representasi data dari tabel database.

## Prasyarat

Sebelum memulai, pastikan Anda telah menginstal:

-   [Go](https://go.dev/doc/install) (versi 1.25.0 atau lebih baru)
-   [PostgreSQL](https://www.postgresql.org/download/)
-   [Goose](https://github.com/pressly/goose) (untuk migrasi database)

## Instalasi & Konfigurasi

1.  **Clone Repositori**
    ```sh
    git clone https://github.com/FalahFaj/Api_funcpro
    cd Api_funcpro
    ```

2.  **Install Dependensi**
    Jalankan perintah berikut untuk mengunduh semua modul yang dibutuhkan.
    ```sh
    go mod tidy
    ```

3.  **Install Goose (Migration Tool)**
    ```sh
    go install github.com/pressly/goose/v3/cmd/goose@latest
    ```

4.  **Konfigurasi Environment**
    Buat file `.env` di direktori root proyek. Anda bisa menyalin dari file `.env.example` yang telah disediakan.
    ```sh
    # Contoh isi file .env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=secret
    DB_NAME=db_pertanian
    DB_SSL_MODE=disable
    ```
    Sesuaikan nilainya dengan konfigurasi database PostgreSQL Anda.

5.  **Setup Database & Migrasi**
    Pastikan server PostgreSQL Anda berjalan dan buat database baru sesuai dengan `DB_NAME` di file `.env`. Kemudian, jalankan migrasi untuk membuat semua tabel yang diperlukan.
    ```sh
    .\migrate.bat up
    ```

## Menjalankan Aplikasi

Setelah semua konfigurasi selesai, jalankan server dengan perintah berikut:

```sh
go run .
```

Server akan berjalan di `http://localhost:8080`.

---
*Dibuat oleh Kelompok 12 - Pemrograman Fungsional.*
