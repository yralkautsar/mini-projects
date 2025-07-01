// mini-projects/todolist-app/todolist-app.go
package todolist_app // Mendeklarasikan package ini sebagai 'todolist_app'

import (
	"bufio" // Untuk membaca input dari pengguna
	"fmt"   // Untuk mencetak output ke terminal
	// Untuk akses ke input/output standar (seperti os.Stdin)
	"strconv" // Untuk konversi string ke tipe numerik
	"strings" // Untuk manipulasi string (misalnya, TrimSpace)
)

// Task adalah struktur data untuk menyimpan detail setiap tugas.
// Huruf awal kapital (Task) menandakan struct ini bisa diakses dari luar package jika diperlukan,
// namun dalam konteks ini, kita hanya akan menggunakannya secara internal.
type Task struct {
	Title    string // Judul tugas
	Complete bool   // Status apakah tugas sudah selesai (true/false)
}

// tasks adalah map (peta) yang menyimpan semua tugas.
// Kunci (key) adalah ID tugas (int), dan nilai (value) adalah objek Task.
// Variabel ini bersifat internal untuk package todolist_app (huruf awal kecil).
var tasks = make(map[int]Task)

// nextID adalah counter untuk ID tugas yang akan datang.
// Variabel ini juga internal untuk package todolist_app.
var nextID = 1

// addTask menambahkan tugas baru ke dalam daftar.
// Fungsi ini internal (huruf awal kecil) karena dipanggil dari RunTodoListCLI.
func addTask(reader *bufio.Reader) {
	fmt.Print("Masukkan nama tugas: ")
	title, _ := reader.ReadString('\n') // Baca judul tugas dari user
	title = strings.TrimSpace(title)    // Hapus spasi dan newline

	// Cek jika judul kosong
	if title == "" {
		fmt.Println("Judul tugas tidak boleh kosong.")
		return
	}

	tasks[nextID] = Task{Title: title, Complete: false} // Tambahkan tugas baru ke map 'tasks'
	fmt.Printf("Tugas '%s' ditambahkan dengan ID %d\n", title, nextID)
	nextID++ // Naikkan ID untuk tugas berikutnya
}

// viewAllTasks menampilkan semua tugas yang ada dalam daftar.
// Fungsi ini internal (huruf awal kecil).
func viewAllTasks() {
	if len(tasks) == 0 { // Cek jika map 'tasks' kosong
		fmt.Println("Belum ada tugas.")
		return
	}

	fmt.Println("\nDaftar Tugas:")
	for id, task := range tasks { // Iterasi melalui setiap tugas di map
		status := "❌" // Status default untuk tugas yang belum selesai
		if task.Complete {
			status = "✅" // Ganti status jika tugas sudah selesai
		}
		fmt.Printf("%d. %s [%s]\n", id, task.Title, status) // Tampilkan ID, judul, dan status tugas
	}
}

// markTaskComplete menandai tugas sebagai selesai berdasarkan ID.
// Fungsi ini internal (huruf awal kecil).
func markTaskComplete(reader *bufio.Reader) {
	fmt.Print("Masukkan ID tugas yang sudah selesai: ")
	idInput, _ := reader.ReadString('\n') // Baca ID tugas dari user
	idInput = strings.TrimSpace(idInput)
	id, err := strconv.Atoi(idInput) // Konversi string ID ke integer
	if err != nil {
		fmt.Println("ID tidak valid. Masukkan angka.")
		return
	}

	task, exists := tasks[id] // Cari tugas berdasarkan ID di map
	if !exists {              // Jika tugas tidak ditemukan
		fmt.Println("Tugas dengan ID tersebut tidak ditemukan.")
		return
	}

	if task.Complete { // Jika tugas sudah selesai
		fmt.Printf("Tugas ID %d ('%s') sudah selesai.\n", id, task.Title)
		return
	}

	task.Complete = true // Tandai tugas sebagai selesai
	tasks[id] = task     // Update status di map 'tasks'
	fmt.Printf("Tugas ID %d ('%s') telah ditandai selesai.\n", id, task.Title)
}

// deleteTask menghapus tugas dari daftar berdasarkan ID.
// Fungsi ini internal (huruf awal kecil).
func deleteTask(reader *bufio.Reader) {
	fmt.Print("Masukkan ID tugas yang ingin dihapus: ")
	idInput, _ := reader.ReadString('\n') // Baca ID tugas dari user
	idInput = strings.TrimSpace(idInput)
	id, err := strconv.Atoi(idInput) // Konversi string ID ke integer
	if err != nil {
		fmt.Println("ID tidak valid. Masukkan angka.")
		return
	}

	_, exists := tasks[id] // Cek apakah tugas dengan ID tersebut ada
	if !exists {           // Jika tidak ditemukan
		fmt.Println("Tugas tidak ditemukan.")
		return
	}

	delete(tasks, id) // Hapus tugas dari map 'tasks'
	fmt.Printf("Tugas ID %d telah dihapus.\n", id)
}

// RunTodoListCLI adalah fungsi utama yang menjalankan aplikasi To-Do List CLI.
// Fungsi ini diekspor (huruf awal kapital 'R') sehingga bisa dipanggil dari package 'main'.
// Parameter 'reader' diperlukan untuk membaca input dari pengguna,
// agar fungsi ini bisa diintegrasikan dengan sistem input di main.go.
func RunTodoListCLI(reader *bufio.Reader) {
	fmt.Println("\n== To-Do List CLI ==")

	for { // Loop utama untuk menampilkan menu dan menerima pilihan pengguna
		// Menampilkan pilihan menu
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Lihat Semua Tugas")
		fmt.Println("3. Tandai Tugas Selesai")
		fmt.Println("4. Hapus Tugas")
		fmt.Println("5. Keluar")

		fmt.Print("Pilih menu (1-5): ")
		input, _ := reader.ReadString('\n') // Baca input pilihan user
		input = strings.TrimSpace(input)    // Hilangkan spasi atau newline
		choice, err := strconv.Atoi(input)  // Konversi string pilihan ke integer
		if err != nil {
			fmt.Println("Masukkan angka yang valid!")
			continue // Lanjut ke iterasi berikutnya (tampilkan menu lagi)
		}

		switch choice { // Memproses pilihan user
		case 1:
			addTask(reader) // Panggil fungsi untuk menambahkan tugas
		case 2:
			viewAllTasks() // Panggil fungsi untuk melihat semua tugas
		case 3:
			markTaskComplete(reader) // Panggil fungsi untuk menandai tugas selesai
		case 4:
			deleteTask(reader) // Panggil fungsi untuk menghapus tugas
		case 5:
			fmt.Println("Keluar dari To-Do List. Sampai jumpa!")
			return // Keluar dari fungsi RunTodoListCLI
		default:
			fmt.Println("Pilihan tidak tersedia. Silakan pilih antara 1-5.")
		}
	}
}
