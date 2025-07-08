// mini-projects/main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// Tidak perlu import log dan net/http di sini lagi karena hanya RunProductAPICLI yang membutuhkannya.

	"mini-projects/bookstore"
	calculator_app "mini-projects/calculator-app"
	contact_manager_app "mini-projects/contact-app"
	parallel_downloader_app "mini-projects/downloader-app"
	"mini-projects/product_service"
	todolist_app "mini-projects/todolist-app"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- My Mini-Projects ---")
		fmt.Println("Choose a Mini Project:")
		fmt.Println("1. Calculator")
		fmt.Println("2. To-Do List")
		fmt.Println("3. Contact Manager")
		fmt.Println("4. Parallel File Downloader")
		fmt.Println("5. Book CRUD App (JSON)")

		// Opsi Start/Stop Server API akan dinamis
		if product_service.IsProductAPIRunning() {
			fmt.Println("6. Stop Product API Server") // Tampilkan opsi stop jika server jalan
		} else {
			fmt.Println("6. Start Product API Server") // Tampilkan opsi start jika server mati
		}

		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			calculator_app.RunCalculatorCLI(reader)
		case "2":
			todolist_app.RunTodoListCLI(reader)
		case "3":
			contact_manager_app.RunContactManagerCLI(reader)
		case "4":
			parallel_downloader_app.RunParallelDownloaderCLI(reader)
		case "5":
			bookstore.RunBookCRUDCLI(reader)
		case "6":
			if product_service.IsProductAPIRunning() {
				product_service.StopProductAPIServer() // Panggil fungsi stop jika server jalan
			} else {
				product_service.RunProductAPICLI() // Panggil fungsi start jika server mati
			}
		case "7":
			fmt.Println("Thank you for using Mini-Projects! Sayonara!")
			// Pastikan server API dihentikan dengan graceful saat keluar aplikasi utama
			if product_service.IsProductAPIRunning() {
				product_service.StopProductAPIServer()
			}
			return
		default:
			fmt.Println("Invalid option. Please choose a number from the menu.")
		}
	}
}
