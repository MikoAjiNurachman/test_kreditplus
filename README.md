# test_kreditplus
# Aplikasi Kredit Plus

Aplikasi Go untuk mengelola pengajuan dan perhitungan kredit, dengan dukungan database PostgreSQL.

## Fitur

* **Pengajuan Kredit:** Konsumen dapat mengajukan kredit dengan memberikan informasi yang diperlukan.
* **Perhitungan Kredit:** Sistem menghitung jumlah cicilan dan bunga berdasarkan jumlah pinjaman, tenor, dan suku bunga.
* **Manajemen Data Konsumen:** Menyimpan dan mengelola data konsumen, termasuk limit kredit dan informasi pribadi.
* **Manajemen Transaksi:** Merekam dan melacak transaksi kredit yang telah disetujui.

## Prasyarat

* **Docker:** Pastikan Anda telah menginstal Docker di mesin Anda. Anda dapat mengunduhnya dari [https://www.docker.com/](https://www.docker.com/).
* **Docker Compose:** Pastikan Anda telah menginstal Docker Compose. Anda dapat mengunduhnya dari [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/).

## Instalasi

1. Clone repositori ini ke local mesin Anda:

   ```bash
   git clone [[URL yang tidak valid dihapus]](https://github.com/MikoAjiNurachman/test_kreditplus)
   ```
2. Buka terminal di direktori proyek dan jalankan perintah berikut untuk menjalankan aplikasi dan database PostgreSQL menggunakan Docker Compose:
   ```bash
   docker-compose up -d
   ```
3. Aplikasi akan berjalan di http://localhost:8045.
   
