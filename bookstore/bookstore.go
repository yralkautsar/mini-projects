// mini-projects/bookstore/bookstore.go
package bookstore // Deklarasi package ini sebagai 'bookstore'

import (
	"bufio"         // Untuk membaca input dari pengguna (di RunBookCRUDCLI)
	"encoding/json" // Untuk mengkodekan (encode) dan mendekodekan (decode) data JSON
	"fmt"           // Untuk fungsi input/output
	"io/ioutil"     // Untuk membaca dan menulis file secara keseluruhan (ioutil.ReadFile/WriteFile)
	"os"            // Untuk operasi sistem file (misalnya mengecek keberadaan file)
	"strconv"       // Untuk konversi string ke tipe numerik dan sebaliknya
	"strings"       // Untuk manipulasi string (misalnya menghilangkan spasi)
)

// Book adalah struct yang merepresentasikan sebuah buku.
// Tag `json:"field_name"` menentukan bagaimana field ini akan dinamai dalam JSON.
// Huruf awal kapital pada field (ID, Title, dll.) membuat mereka diekspor.
type Book struct {
	ID     int    `json:"id"`     // ID unik untuk buku
	Title  string `json:"title"`  // Judul buku
	Author string `json:"author"` // Penulis buku
	Year   int    `json:"year"`   // Tahun terbit
}

// bookDB adalah slice (array dinamis) yang akan menyimpan semua data buku dalam memori.
// Variabel ini bersifat internal (tidak diekspor) untuk package bookstore (huruf awal kecil).
var bookDB []Book

// jsonFilePath adalah path ke file JSON tempat data buku akan disimpan.
// Ini adalah konstanta internal package.
const jsonFilePath = "books.json"

// --- Fungsi Bantuan untuk File I/O JSON (Internal) ---

// loadBooksFromJsonFile memuat (membaca) data buku dari file JSON ke 'bookDB'.
// Fungsi ini bersifat internal (huruf awal kecil) karena hanya dipanggil dari dalam package bookstore.
func loadBooksFromJsonFile() error {
	// Cek apakah file JSON ada. Jika tidak ada, berarti ini pertama kali dijalankan,
	// jadi kita tidak perlu memuat apa-apa dan kembali tanpa error.
	if _, err := os.Stat(jsonFilePath); os.IsNotExist(err) {
		fmt.Println("File 'books.json' tidak ditemukan. Membuat database buku kosong.")
		bookDB = []Book{} // Inisialisasi slice kosong
		return nil
	}

	// Membaca seluruh konten file JSON.
	data, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		return fmt.Errorf("gagal membaca file JSON: %w", err)
	}

	// Jika file kosong, inisialisasi bookDB kosong.
	if len(data) == 0 {
		bookDB = []Book{}
		fmt.Println("File 'books.json' kosong. Membuat database buku kosong.")
		return nil
	}

	// Mendecode data JSON (byte slice) ke slice Book (bookDB).
	// Tanda '&' penting karena kita memberikan pointer ke variabel yang akan diisi.
	if err := json.Unmarshal(data, &bookDB); err != nil {
		return fmt.Errorf("gagal mendekode JSON dari file: %w", err)
	}

	fmt.Println("Data buku berhasil dimuat dari 'books.json'.")
	return nil
}

// saveBooksToJsonFile menyimpan (menulis) data buku dari 'bookDB' ke file JSON.
// Fungsi ini bersifat internal (huruf awal kecil) karena hanya dipanggil dari dalam package bookstore.
func saveBooksToJsonFile() error {
	// Mengkodekan (encode) slice 'bookDB' menjadi format JSON.
	// json.MarshalIndent menghasilkan output JSON yang rapi (pretty-printed)
	// dengan indentasi (prefix="", indent="  ").
	data, err := json.MarshalIndent(bookDB, "", "  ")
	if err != nil {
		return fmt.Errorf("gagal mengkodekan data ke JSON: %w", err)
	}

	// Menulis data JSON (byte slice) ke file.
	// os.WriteFile akan menulis/menimpa file dengan data baru.
	// 0644 adalah permission file (bisa dibaca dan ditulis oleh pemilik, hanya dibaca oleh grup/lainnya).
	if err := os.WriteFile(jsonFilePath, data, 0644); err != nil {
		return fmt.Errorf("gagal menulis data JSON ke file: %w", err)
	}

	fmt.Println("Data buku berhasil disimpan ke 'books.json'.")
	return nil
}

// --- Fungsi-fungsi CRUD (Diekspor untuk digunakan oleh main.go) ---

// GetNextID menemukan ID unik berikutnya untuk buku baru.
// Fungsi ini diekspor agar bisa dipanggil dari main.go jika diperlukan,
// atau tetap sebagai pembantu internal. Untuk demo ini, kita ekspor.
func GetNextID() int {
	maxID := 0
	for _, book := range bookDB {
		if book.ID > maxID {
			maxID = book.ID
		}
	}
	return maxID + 1
}

// AddBook (CREATE) menambahkan buku baru ke bookDB dan menyimpannya.
// Fungsi ini diekspor (huruf awal kapital 'A') agar bisa dipanggil dari main.go.
func AddBook(title, author string, year int) error {
	newBook := Book{
		ID:     GetNextID(), // Dapatkan ID unik (memanggil fungsi yang diekspor GetNextID)
		Title:  title,
		Author: author,
		Year:   year,
	}
	bookDB = append(bookDB, newBook) // Tambahkan buku ke memori

	if err := saveBooksToJsonFile(); err != nil { // Simpan perubahan ke file JSON (memanggil fungsi internal)
		return fmt.Errorf("error menyimpan buku: %w", err)
	}
	fmt.Printf("Buku '%s' berhasil ditambahkan (ID: %d).\n", newBook.Title, newBook.ID)
	return nil
}

// ListBooks (READ) menampilkan semua buku yang ada di bookDB.
// Fungsi ini diekspor (huruf awal kapital 'L') agar bisa dipanggil dari main.go.
func ListBooks() {
	fmt.Println("\n--- Daftar Buku ---")
	if len(bookDB) == 0 {
		fmt.Println("Tidak ada buku dalam daftar.")
		return
	}
	for _, book := range bookDB {
		fmt.Printf("ID: %d, Judul: \"%s\", Penulis: \"%s\", Tahun: %d\n",
			book.ID, book.Title, book.Author, book.Year)
	}
}

// UpdateBook (UPDATE) memperbarui detail buku berdasarkan ID dan menyimpannya.
// Fungsi ini diekspor (huruf awal kapital 'U') agar bisa dipanggil dari main.go.
func UpdateBook(id int, newTitle, newAuthor string, newYear int) error {
	foundIndex := -1
	for i, book := range bookDB {
		if book.ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return fmt.Errorf("buku dengan ID %d tidak ditemukan", id)
	}

	// Perbarui hanya jika ada nilai baru yang diberikan (tidak kosong/nol)
	// Kita berasumsi bahwa string kosong berarti tidak ada perubahan.
	// Untuk int, 0 adalah zero value, jadi hati-hati jika 0 adalah nilai valid yang ingin di-set.
	if newTitle != "" {
		bookDB[foundIndex].Title = newTitle
	}
	if newAuthor != "" {
		bookDB[foundIndex].Author = newAuthor
	}
	if newYear != 0 {
		bookDB[foundIndex].Year = newYear
	}

	if err := saveBooksToJsonFile(); err != nil { // Simpan perubahan ke file JSON
		return fmt.Errorf("error menyimpan perubahan: %w", err)
	}
	fmt.Printf("Buku ID %d berhasil diperbarui.\n", id)
	return nil
}

// DeleteBook (DELETE) menghapus buku berdasarkan ID dan menyimpannya.
// Fungsi ini diekspor (huruf awal kapital 'D') agar bisa dipanggil dari main.go.
func DeleteBook(id int) error {
	found := false
	newBookDB := []Book{} // Buat slice baru untuk menyimpan buku yang tidak dihapus
	for _, book := range bookDB {
		if book.ID == id {
			found = true
		} else {
			newBookDB = append(newBookDB, book) // Tambahkan buku ke slice baru jika bukan yang dihapus
		}
	}

	if !found {
		return fmt.Errorf("buku dengan ID %d tidak ditemukan", id)
	}

	bookDB = newBookDB // Ganti bookDB dengan slice yang sudah difilter

	if err := saveBooksToJsonFile(); err != nil { // Simpan perubahan ke file JSON
		return fmt.Errorf("error menyimpan perubahan: %w", err)
	}
	fmt.Printf("Buku ID %d berhasil dihapus.\n", id)
	return nil
}

// RunBookCRUDCLI adalah fungsi utama yang menjalankan aplikasi CRUD Buku CLI.
// Fungsi ini diekspor (huruf awal kapital 'R') sehingga bisa dipanggil dari package 'main'.
// Parameter 'reader' diperlukan untuk membaca input dari pengguna.
func RunBookCRUDCLI(reader *bufio.Reader) {
	fmt.Println("\n== Aplikasi CRUD Buku dengan JSON ==")

	// Inisialisasi sistem buku dengan memuat data dari file JSON.
	// Ini adalah panggilan ke fungsi internal 'loadBooksFromJsonFile'.
	if err := loadBooksFromJsonFile(); err != nil {
		fmt.Printf("Error saat menginisialisasi database buku: %v\n", err)
		return // Keluar dari fungsi jika ada error fatal saat inisialisasi
	}

	for { // Loop utama menu aplikasi CRUD buku
		fmt.Println("\nMenu Buku:")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Lihat Semua Buku")
		fmt.Println("3. Perbarui Buku")
		fmt.Println("4. Hapus Buku")
		fmt.Println("5. Kembali ke Menu Utama") // Opsi untuk kembali ke menu utama My Super Go App
		fmt.Print("Pilih opsi: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Println("\n--- Tambah Buku Baru ---")
			fmt.Print("Judul: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Penulis: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)

			fmt.Print("Tahun Terbit: ")
			yearStr, _ := reader.ReadString('\n')
			year, err := strconv.Atoi(strings.TrimSpace(yearStr))
			if err != nil {
				fmt.Println("Tahun terbit tidak valid. Pembatalan.")
				continue
			}
			// Memanggil fungsi 'AddBook' yang diekspor
			if err := AddBook(title, author, year); err != nil {
				fmt.Printf("Error menambah buku: %v\n", err)
			}
		case "2":
			// Memanggil fungsi 'ListBooks' yang diekspor
			ListBooks()
		case "3":
			fmt.Println("\n--- Perbarui Buku ---")
			fmt.Print("Masukkan ID buku yang ingin diperbarui: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("ID tidak valid. Pembatalan.")
				continue
			}

			fmt.Print("Judul Baru (biarkan kosong jika tidak berubah): ")
			newTitle, _ := reader.ReadString('\n')
			newTitle = strings.TrimSpace(newTitle)

			fmt.Print("Penulis Baru (biarkan kosong jika tidak berubah): ")
			newAuthor, _ := reader.ReadString('\n')
			newAuthor = strings.TrimSpace(newAuthor)

			fmt.Print("Tahun Terbit Baru (biarkan kosong jika tidak berubah): ")
			newYearStr, _ := reader.ReadString('\n')
			newYear := 0 // Default 0 jika tidak diisi atau tidak valid
			if strings.TrimSpace(newYearStr) != "" {
				parsedYear, err := strconv.Atoi(strings.TrimSpace(newYearStr))
				if err != nil {
					fmt.Println("Tahun terbit baru tidak valid. Perubahan tahun dibatalkan.")
				} else {
					newYear = parsedYear
				}
			}
			// Memanggil fungsi 'UpdateBook' yang diekspor
			if err := UpdateBook(id, newTitle, newAuthor, newYear); err != nil {
				fmt.Printf("Error memperbarui buku: %v\n", err)
			}
		case "4":
			fmt.Println("\n--- Hapus Buku ---")
			fmt.Print("Masukkan ID buku yang ingin dihapus: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("ID tidak valid. Pembatalan.")
				continue
			}
			// Memanggil fungsi 'DeleteBook' yang diekspor
			if err := DeleteBook(id); err != nil {
				fmt.Printf("Error menghapus buku: %v\n", err)
			}
		case "5":
			fmt.Println("Kembali ke Menu Utama.")
			return // Keluar dari fungsi ini, kembali ke loop main() di main.go
		default:
			fmt.Println("Opsi tidak valid. Silakan pilih antara 1-5.")
		}
	}
}
