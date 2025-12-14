# Proyek API E-Commerce Pertanian (FuncPro - Kelompok 11)

Ini adalah projek API yang simpel aja, penting selesai

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

## Instalasi & Konfigurasi

1.  **Clone Repositori**
    ```sh
    git clone https://github.com/FalahFaj/Api_funcpro
    cd Api_funcpro
    ```
    jalankan di terminal
    ```
    git checkout -b nama_kalian
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
    Sesuaikan nilai di file `.env` sesuai konfigurasi postgre kalian.

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

Server akan berjalan di `http://localhost:8030`.

---
*Dibuat oleh Kelompok 11 - Pemrograman Fungsional.*
