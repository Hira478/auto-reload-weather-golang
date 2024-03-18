# Proyek Auto Reload Cuaca

Proyek ini merupakan sebuah microservice menggunakan Go yang menghasilkan nilai acak untuk air dan angin setiap 15 detik, memperbarui sebuah file JSON (`status.json`), dan menyajikan halaman web yang menampilkan status berdasarkan kriteria yang telah ditentukan sebelumnya.

## Fitur

- Menghasilkan nilai acak untuk air dan angin setiap 15 detik.
- Memperbarui file JSON (`status.json`) dengan nilai terbaru.
- Menyajikan halaman web yang menampilkan status berdasarkan kriteria yang telah ditentukan sebelumnya:
  - Jika air berada di bawah 5 dan angin di bawah 6, statusnya adalah "Aman".
  - Jika air berada antara 6 dan 8 dan angin berada antara 6 dan 15, statusnya adalah "Siaga".
  - Jika air berada di atas 8 atau angin berada di atas 15, statusnya adalah "Bahaya".

## Penggunaan

1. Pastikan Anda memiliki Go terinstal di sistem Anda.
2. Klon repositori ini.
3. Buka direktori proyek.
4. Jalankan perintah berikut untuk memulai server:
   `go run main.go`
5. Buka peramban web Anda dan buka http://localhost:8080 untuk melihat halaman status.

## Struktur Data

- main.go: Berisi logika server dan kode untuk menghasilkan dan melayani file status.json.
- static/index.html: Berkas HTML untuk menampilkan halaman status.
