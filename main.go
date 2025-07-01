// main.go
package main // Deklarasi package utama yang bisa dieksekusi oleh Go

import (
	"bufio"   // Untuk membaca input dari user (keyboard)
	"fmt"     // Untuk mencetak output ke konsol
	"os"      // Untuk interaksi dengan sistem operasi, seperti membaca dari input standar (os.Stdin)
	"strings" // Untuk operasi string seperti menghapus spasi atau karakter newline

	// Mengimpor package-package mini proyek yang sudah dimodularisasi.
	// Path ini harus sesuai dengan nama module yang kamu definisikan di 'go.mod'
	// ditambah dengan nama folder dari masing-masing package.
	// Contoh: 'mini-projects/calculator-app'
	calculator_app "mini-projects/calculator-app"          // Impor package kalkulator
	contact_manager_app "mini-projects/contact-app"        // Impor package contact manager
	parallel_downloader_app "mini-projects/downloader-app" // Impor package parallel downloader
	todolist_app "mini-projects/todolist-app"              // Impor package to-do list
)

func main() {
	// Membuat objek 'reader' untuk membaca input dari keyboard.
	// Ini akan digunakan oleh semua mini proyek yang memerlukan input dari user.
	reader := bufio.NewReader(os.Stdin)

	for { // Loop utama aplikasi, akan terus menampilkan menu sampai user memilih untuk keluar.
		fmt.Println("\n--- Mini Projects by YOGA ---")
		fmt.Println("Choose a Mini Project:")
		fmt.Println("1. Calculator")
		fmt.Println("2. To-Do List")
		fmt.Println("3. Contact Manager")
		fmt.Println("4. Parallel File Downloader") // Opsi untuk menjalankan Downloader
		fmt.Println("5. Exit")                     // Opsi untuk keluar dari aplikasi gabungan
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n') // Membaca seluruh baris input sampai user menekan Enter.
		input = strings.TrimSpace(input)    // Menghapus spasi di awal/akhir dan karakter newline dari input.

		switch input { // Memeriksa input user dan mengarahkan ke mini proyek yang dipilih.
		case "1":
			// Memanggil fungsi utama dari package 'calculator-app'.
			// Fungsi ini akan mengambil alih kontrol CLI sampai user keluar dari kalkulator.
			calculator_app.RunCalculatorCLI(reader)
		case "2":
			// Memanggil fungsi utama dari package 'todolist-app'.
			todolist_app.RunTodoListCLI(reader)
		case "3":
			// Memanggil fungsi utama dari package 'contact-app'.
			contact_manager_app.RunContactManagerCLI(reader)
		case "4":
			// Memanggil fungsi utama dari package 'downloader-app'.
			parallel_downloader_app.RunParallelDownloaderCLI(reader)
		case "5":
			fmt.Println("Thank you for using My Super Go App! Goodbye.")
			return // Keluar dari fungsi main, yang akan menghentikan eksekusi program.
		default:
			fmt.Println("Invalid option. Please choose a number from the menu.")
		}
	}
}
