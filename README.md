# **Go Mini-Projects**

Kumpulan aplikasi kecil mandiri (mini-proyek) yang dibangun dengan Go (Golang). Repository ini berfungsi sebagai tempat belajar berbagai fungsionalitas Go, mulai dari aplikasi CLI dasar hingga API web sederhana.

## **🌟 Fitur**

Proyek ini menyediakan antarmuka baris perintah (CLI) pusat untuk mengakses beberapa aplikasi mini:

* **Kalkulator CLI:** Sebuah kalkulator baris perintah dasar untuk melakukan operasi aritmatika.  
  * Mendukung operasi: penjumlahan (+), pengurangan (-), perkalian (\*), pembagian (/), dan pangkat (^).  
  * Mampu memproses ekspresi berantai (misal: 5 \+ 3 \- 1).  
  * Menyediakan riwayat perhitungan yang dapat dilihat dengan mengetik history.  
  * Dapat keluar dari kalkulator dengan mengetik exit.  
* **Daftar Tugas (To-Do List) CLI:** Kelola tugas Anda langsung dari terminal.  
  * **Tambah Tugas:** Menambahkan tugas baru ke daftar.  
  * **Lihat Semua Tugas:** Menampilkan semua tugas beserta status selesai/belum selesai.  
  * **Tandai Tugas Selesai:** Menandai tugas sebagai selesai berdasarkan ID-nya.  
  * **Hapus Tugas:** Menghapus tugas dari daftar berdasarkan ID-nya.  
  * Data tugas disimpan dalam memori selama aplikasi berjalan.  
* **Manajer Kontak CLI:** Aplikasi sederhana berbasis baris perintah untuk menyimpan dan mengelola kontak.  
  * **Tambah Kontak:** Menambahkan nama dan nomor telepon kontak baru.  
  * **Lihat Semua Kontak:** Menampilkan daftar semua kontak yang tersimpan.  
  * **Hapus Kontak:** Menghapus kontak dari daftar berdasarkan namanya.  
  * Data kontak disimpan dalam memori selama aplikasi berjalan.  
* **Pengunduh File Paralel CLI:** Unduh file besar secara efisien dengan memecahnya menjadi beberapa bagian dan mengunduhnya secara bersamaan menggunakan Goroutine dan Channel Go.  
  * Mendukung pengunduhan file dari URL.  
  * Memungkinkan pengguna menentukan jumlah bagian paralel untuk pengunduhan.  
  * Secara otomatis menggabungkan bagian-bagian yang diunduh menjadi satu file lengkap.  
  * Menangani error saat pengunduhan dan membersihkan file sementara jika terjadi kegagalan.  
* **Aplikasi CRUD Buku (JSON) CLI:** Lakukan operasi Buat, Baca, Perbarui, dan Hapus (CRUD) untuk buku.  
  * **Tambah Buku:** Menambahkan buku baru dengan judul, penulis, dan tahun terbit.  
  * **Lihat Semua Buku:** Menampilkan daftar semua buku yang tersimpan.  
  * **Perbarui Buku:** Memperbarui detail buku (judul, penulis, tahun) berdasarkan ID buku.  
  * **Hapus Buku:** Menghapus buku dari daftar berdasarkan ID-nya.  
  * Data buku disimpan secara persisten dalam file books.json, sehingga data tidak hilang saat aplikasi ditutup.  
* **Server API Produk:** Layanan API RESTful yang memungkinkan operasi CRUD (Create, Read, Update, Delete) untuk data produk.  
  * Berjalan sebagai server HTTP terpisah yang dapat dimulai dan dihentikan dari menu utama aplikasi CLI.  
  * Data produk disimpan secara persisten dalam file products.json.  
  * Mendukung endpoint RESTful untuk mengelola koleksi produk dan produk individual.

## **🚀 Memulai**

Instruksi ini akan membantu Anda mendapatkan salinan proyek dan menjalankannya di mesin lokal Anda untuk tujuan pengembangan dan pengujian.

### **Prasyarat**

* [Go (Golang)](https://golang.org/doc/install) terinstal di sistem Anda (versi 1.22 atau lebih tinggi direkomendasikan).

### **Instalasi**

1. **Kloning repositori:**  
   git clone https://github.com/your-username/your-repo-name.git  
   cd your-repo-name

   *(Ganti your-username/your-repo-name dengan username dan nama repositori Anda)*  
2. **Inisialisasi modul Go dan unduh dependensi:**  
   go mod tidy

### **Menjalankan Aplikasi**

Untuk menjalankan aplikasi utama dan mengakses mini-proyek:

go run main.go

Ini akan meluncurkan menu utama di terminal Anda, memungkinkan Anda memilih mini-proyek mana yang akan dijalankan.

## **💡 Penggunaan**

Setelah aplikasi berjalan, Anda akan melihat menu seperti ini:

\--- My Super Go App \---  
Choose a Mini Project:  
1\. Calculator  
2\. To-Do List  
3\. Contact Manager  
4\. Parallel File Downloader  
5\. Book CRUD App (JSON)  
6\. Start Product API Server (or Stop Product API Server if running)  
7\. Exit  
Enter your choice:

* **Pilih nomor (1-6)** untuk menjalankan mini-proyek yang sesuai.  
* **Opsi 6** akan mengaktifkan/menonaktifkan server API Produk (mulai jika berhenti, berhenti jika berjalan).  
* **Opsi 7** akan keluar dari aplikasi.

### **Penggunaan Kalkulator CLI**

Saat Anda memilih opsi "1. Calculator" dari menu utama, Anda akan masuk ke mode kalkulator:

\== Kalkulator CLI Sederhana \==  
Ketik 'exit' untuk keluar, 'history' untuk melihat riwayat.  
Operator yang didukung: \+  \-  \* /  ^

Masukkan ekspresi (contoh: 5 \+ 3 \- 1):

* Masukkan ekspresi matematika (misal: 10 \* 5 \+ 2).  
* Ketik history untuk melihat riwayat perhitungan sebelumnya.  
* Ketik exit untuk kembali ke menu utama.

### **Penggunaan To-Do List CLI**

Saat Anda memilih opsi "2. To-Do List" dari menu utama, Anda akan masuk ke menu To-Do List:

\== To-Do List CLI \==

Menu:  
1\. Tambah Tugas  
2\. Lihat Semua Tugas  
3\. Tandai Tugas Selesai  
4\. Hapus Tugas  
5\. Keluar  
Pilih menu (1-5):

* Pilih 1 untuk menambahkan tugas baru.  
* Pilih 2 untuk melihat semua tugas yang ada.  
* Pilih 3 untuk menandai tugas sebagai selesai dengan memasukkan ID tugas.  
* Pilih 4 untuk menghapus tugas dengan memasukkan ID tugas.  
* Pilih 5 untuk keluar dari aplikasi To-Do List dan kembali ke menu utama.

### **Penggunaan Manajer Kontak CLI**

Saat Anda memilih opsi "3. Contact Manager" dari menu utama, Anda akan masuk ke menu Manajer Kontak:

\== Contact Manager CLI \==

Menu:  
1\. Add Contact  
2\. View All Contacts  
3\. Delete Contact  
4\. Back to Main Menu  
Choose option (1-4):

* Pilih 1 untuk menambahkan kontak baru (Anda akan diminta untuk memasukkan Nama dan Nomor Telepon).  
* Pilih 2 untuk melihat daftar semua kontak yang tersimpan.  
* Pilih 3 untuk menghapus kontak dengan memasukkan nama kontak yang ingin dihapus.  
* Pilih 4 untuk keluar dari aplikasi Manajer Kontak dan kembali ke menu utama.

### **Penggunaan Pengunduh File Paralel CLI**

Saat Anda memilih opsi "4. Parallel File Downloader" dari menu utama, Anda akan diminta untuk memasukkan detail pengunduhan:

\--- Go Parallel File Downloader \---  
Masukkan URL file (contoh: https://speed.cloudflare.com/\_\_down?bytes=10000000):  
Masukkan nama file output (contoh: downloaded\_10MB.bin):  
Masukkan jumlah bagian paralel (contoh: 4):

* Masukkan URL lengkap dari file yang ingin Anda unduh.  
* Berikan nama untuk file yang akan disimpan setelah pengunduhan selesai.  
* Tentukan berapa banyak bagian paralel yang ingin Anda gunakan untuk mengunduh file. Semakin banyak bagian, semakin banyak Goroutine yang akan digunakan.

Aplikasi akan menampilkan progres pengunduhan setiap bagian dan kemudian menggabungkan semua bagian menjadi satu file setelah selesai.

### **Penggunaan Aplikasi CRUD Buku (JSON) CLI**

Saat Anda memilih opsi "5. Book CRUD App (JSON)" dari menu utama, Anda akan masuk ke menu manajemen buku:

\== Aplikasi CRUD Buku dengan JSON \==

Menu Buku:  
1\. Tambah Buku  
2\. Lihat Semua Buku  
3\. Perbarui Buku  
4\. Hapus Buku  
5\. Kembali ke Menu Utama  
Pilih opsi:

* Pilih 1 untuk menambahkan buku baru (Anda akan diminta untuk memasukkan Judul, Penulis, dan Tahun Terbit).  
* Pilih 2 untuk melihat daftar semua buku yang tersimpan.  
* Pilih 3 untuk memperbarui detail buku (Judul, Penulis, Tahun) dengan memasukkan ID buku yang ingin diperbarui.  
* Pilih 4 untuk menghapus buku dengan memasukkan ID buku yang ingin dihapus.  
* Pilih 5 untuk kembali ke menu utama "My Super Go App".

### **Penggunaan Server API Produk**

Saat Anda memilih opsi "6. Start Product API Server" dari menu utama, server API akan dimulai dan berjalan di http://localhost:8080. Anda akan melihat output di konsol yang menunjukkan server telah dimulai dan endpoint yang tersedia:

\--- REST API Produk Server Go Dimulai \---  
Server API berjalan di http://localhost:8080  
Endpoint API Produk:  
   \[GET\]    /api/products  
   \[POST\]   /api/products  
   \[GET\]    /api/products/{id}  
   \[PUT\]    /api/products/{id}  
   \[DELETE\] /api/products/{id}

Server API siap. Pilih opsi 'Stop Product API Server' di menu untuk kembali.  
Atau tekan Ctrl+C untuk menghentikan seluruh aplikasi.

Anda dapat berinteraksi dengan API ini menggunakan alat seperti curl atau Postman. Berikut adalah beberapa contoh:

**1\. Mendapatkan Semua Produk (GET /api/products)**

curl http://localhost:8080/api/products

Contoh respons:

\[  
  {"id": 1, "name": "Laptop", "price": 1200, "stock": 50},  
  {"id": 2, "name": "Mouse", "price": 25, "stock": 200}  
\]

**2\. Menambahkan Produk Baru (POST /api/products)**

curl \-X POST \-H "Content-Type: application/json" \-d '{"name": "Keyboard", "price": 75, "stock": 150}' http://localhost:8080/api/products

Contoh respons:

{"id": 3, "name": "Keyboard", "price": 75, "stock": 150}

**3\. Mendapatkan Produk Berdasarkan ID (GET /api/products/{id})**

curl http://localhost:8080/api/products/1

Contoh respons:

{"id": 1, "name": "Laptop", "price": 1200, "stock": 50}

**4\. Memperbarui Produk Berdasarkan ID (PUT /api/products/{id})**

curl \-X PUT \-H "Content-Type: application/json" \-d '{"name": "Gaming Laptop", "price": 1300, "stock": 45}' http://localhost:8080/api/products/1

Contoh respons:

{"id": 1, "name": "Gaming Laptop", "price": 1300, "stock": 45}

**5\. Menghapus Produk Berdasarkan ID (DELETE /api/products/{id})**

curl \-X DELETE http://localhost:8080/api/products/2

Respons: (Biasanya tidak ada konten, status 204 No Content)

Untuk menghentikan server API, pilih opsi "6. Stop Product API Server" lagi dari menu utama aplikasi CLI.

## **📁 Struktur Proyek**

Proyek ini diatur ke dalam beberapa paket, dengan main.go bertindak sebagai titik masuk dan orkestrator:

mini-projects/  
├── main.go  
├── bookstore/  
│   └── ... (File aplikasi CRUD Buku)  
├── calculator-app/  
│   └── ... (File aplikasi Kalkulator)  
├── contact-app/  
│   └── ... (File aplikasi Manajer Kontak)  
├── downloader-app/  
│   └── ... (File aplikasi Pengunduh File Paralel)  
├── product\_service/  
│   └── ... (File layanan API Produk)  
├── todolist-app/  
│   └── ... (File aplikasi Daftar Tugas)  
└── go.mod  
└── go.sum

Setiap sub-direktori (bookstore, calculator-app, dll.) berisi kode untuk mini-proyeknya masing-masing.

## **🤝 Berkontribusi**

Kontribusi dipersilakan\! Jika Anda memiliki saran untuk perbaikan, mini-proyek baru, atau perbaikan bug, jangan ragu untuk:

1. Fork repositori.  
2. Buat cabang baru (git checkout \-b feature/AmazingFeature).  
3. Lakukan perubahan Anda.  
4. Commit perubahan Anda (git commit \-m 'Add some AmazingFeature').  
5. Push ke cabang (git push origin feature/AmazingFeature).  
6. Buka Pull Request.

**Dibuat dengan ❤️ oleh \[YOGS / ICS TRAVELGROUP / IT TEAM\]**