// mini-projects/product_service/product_service.go
package product_service

import (
	"context" // Tambahkan import context
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time" // Tambahkan import time
)

// --- Struktur Data (Sama seperti sebelumnya) ---
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock,omitempty"`
}

var products []Product
var nextProductID = 1

const jsonFilePath = "products.json"

// serverInstance menyimpan instance HTTP server agar bisa diakses dan dimatikan.
var serverInstance *http.Server

// --- Fungsi Bantuan untuk File I/O JSON (Sama seperti sebelumnya) ---
func getNextProductID() int {
	currentMaxID := 0
	for _, p := range products {
		if p.ID > currentMaxID {
			currentMaxID = p.ID
		}
	}
	return currentMaxID + 1
}

func loadProductsFromJsonFile() error {
	if _, err := os.Stat(jsonFilePath); os.IsNotExist(err) {
		log.Printf("LOG: File '%s' tidak ditemukan. Membuat database produk kosong.", jsonFilePath)
		products = []Product{}
		return nil
	}
	data, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		return fmt.Errorf("gagal membaca file JSON: %w", err)
	}
	if len(data) == 0 {
		products = []Product{}
		log.Printf("LOG: File '%s' kosong. Membuat database produk kosong.", jsonFilePath)
		return nil
	}
	if err := json.Unmarshal(data, &products); err != nil {
		return fmt.Errorf("gagal mendekode JSON dari file: %w", err)
	}
	log.Printf("LOG: Data produk berhasil dimuat dari '%s'.", jsonFilePath)
	return nil
}

func saveProductsToJsonFile() error {
	data, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return fmt.Errorf("gagal mengkodekan data ke JSON: %w", err)
	}
	if err := os.WriteFile(jsonFilePath, data, 0644); err != nil {
		return fmt.Errorf("gagal menulis data JSON ke file: %w", err)
	}
	log.Printf("LOG: Data produk berhasil disimpan ke '%s'.", jsonFilePath)
	return nil
}

// --- Fungsi Helper untuk Respons API (Sama seperti sebelumnya) ---
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, `{"message": "Internal Server Error"}`, http.StatusInternalServerError)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// --- API Handlers (Sama seperti sebelumnya) ---

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if strings.TrimPrefix(r.URL.Path, "/api/products") != "" && strings.TrimPrefix(r.URL.Path, "/api/products/") != "" {
		respondWithError(w, http.StatusNotFound, "Endpoint tidak ditemukan")
		return
	}
	switch r.Method {
	case "GET":
		respondWithJSON(w, http.StatusOK, products)
		log.Println("LOG: Permintaan GET /api/products berhasil diproses.")
	case "POST":
		var newProduct Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			respondWithError(w, http.StatusBadRequest, "Format JSON permintaan tidak valid")
			log.Printf("Error decoding JSON for POST /api/products: %v", err)
			return
		}
		if newProduct.Name == "" || newProduct.Price <= 0 {
			respondWithError(w, http.StatusBadRequest, "Nama dan Harga produk tidak boleh kosong atau nol")
			return
		}
		newProduct.ID = getNextProductID()
		products = append(products, newProduct)
		if err := saveProductsToJsonFile(); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Gagal menyimpan data produk")
			log.Printf("Error saving products after POST /api/products: %v", err)
			return
		}
		respondWithJSON(w, http.StatusCreated, newProduct)
		log.Printf("LOG: Produk baru ditambahkan: ID %d, Nama '%s'.", newProduct.ID, newProduct.Name)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Metode tidak diizinkan")
		log.Printf("LOG: Metode %s tidak diizinkan untuk /api/products.", r.Method)
	}
}

func productByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID produk tidak valid")
		log.Printf("LOG: ID tidak valid di URL: '%s'", idStr)
		return
	}
	var foundProduct *Product
	foundIndex := -1
	for i, product := range products {
		if product.ID == id {
			foundProduct = &products[i]
			foundIndex = i
			break
		}
	}
	if foundProduct == nil {
		respondWithError(w, http.StatusNotFound, "Produk tidak ditemukan")
		log.Printf("LOG: Produk dengan ID %d tidak ditemukan.", id)
		return
	}
	switch r.Method {
	case "GET":
		respondWithJSON(w, http.StatusOK, foundProduct)
		log.Printf("LOG: Permintaan GET /api/products/%d berhasil diproses.", id)
	case "PUT":
		var updatedProduct Product
		if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
			respondWithError(w, http.StatusBadRequest, "Format JSON permintaan tidak valid")
			log.Printf("Error decoding JSON for PUT /api/products/%d: %v", id, err)
			return
		}
		if updatedProduct.Name != "" {
			foundProduct.Name = updatedProduct.Name
		}
		if updatedProduct.Price != 0 {
			foundProduct.Price = updatedProduct.Price
		}
		foundProduct.Stock = updatedProduct.Stock
		if err := saveProductsToJsonFile(); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Gagal menyimpan perubahan produk")
			log.Printf("Error saving products after PUT /api/products/%d: %v", id, err)
			return
		}
		respondWithJSON(w, http.StatusOK, foundProduct)
		log.Printf("LOG: Produk ID %d berhasil diperbarui.", id)
	case "DELETE":
		newProducts := make([]Product, 0)
		for i, product := range products {
			if i != foundIndex {
				newProducts = append(newProducts, product)
			}
		}
		products = newProducts
		if err := saveProductsToJsonFile(); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Gagal menyimpan perubahan produk (setelah hapus)")
			log.Printf("Error saving products after DELETE /api/products/%d: %v", id, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		log.Printf("LOG: Produk ID %d berhasil dihapus.", id)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Metode tidak diizinkan")
		log.Printf("LOG: Metode %s tidak diizinkan untuk /api/products/{id}.", r.Method)
	}
}

// RunProductAPICLI adalah fungsi yang akan dijalankan ketika opsi API Produk dipilih dari menu CLI.
// Fungsi ini akan menjalankan HTTP server di Goroutine terpisah dan menunggu sinyal stop.
func RunProductAPICLI() {
	// Muat data awal dari JSON saat API diinisialisasi.
	if err := loadProductsFromJsonFile(); err != nil {
		log.Printf("LOG: Error memuat data awal produk dari JSON: %v", err)
		// Tidak pakai log.Fatal di sini agar aplikasi utama tidak mati
		// jika file JSON bermasalah, tapi kita log errornya saja.
	}
	nextProductID = getNextProductID()

	mux := http.NewServeMux() // Membuat router (ServeMux) baru khusus untuk API ini.
	mux.HandleFunc("/api/products", productsHandler)
	mux.HandleFunc("/api/products/", productByIDHandler)

	const apiPort = ":8080" // Port untuk API ini
	// Membuat instance HTTP server
	serverInstance = &http.Server{
		Addr:    apiPort,
		Handler: mux, // Gunakan router yang sudah kita definisikan
	}

	// Channel untuk memberi sinyal bahwa server sudah berhenti
	done := make(chan bool)

	// Jalankan server di Goroutine terpisah.
	// Ini memungkinkan fungsi RunProductAPICLI untuk tidak memblokir dan
	// bisa menampilkan pesan bahwa server berjalan dan menunggu input stop.
	go func() {
		fmt.Printf("\n--- REST API Produk Server Go Dimulai ---\n")
		fmt.Printf("Server API berjalan di http://localhost%s\n", apiPort)
		fmt.Println("Endpoint API Produk:")
		fmt.Println("  [GET]    /api/products")
		fmt.Println("  [POST]   /api/products")
		fmt.Println("  [GET]    /api/products/{id}")
		fmt.Println("  [PUT]    /api/products/{id}")
		fmt.Println("  [DELETE] /api/products/{id}")
		fmt.Println("\nServer API siap. Pilih opsi 'Stop Product API Server' di menu untuk kembali.")
		fmt.Println("Atau tekan Ctrl+C untuk menghentikan seluruh aplikasi.") // Ini akan tetap menghentikan seluruh aplikasi
		// karena Ctrl+C adalah sinyal OS global.
		// Nanti kita akan tambahkan opsi stop di menu.

		if err := serverInstance.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("LOG: Error server API: %v", err) // Server akan mati jika ada error selain ErrServerClosed
		}
		close(done) // Memberi sinyal bahwa Goroutine server sudah selesai.
	}()

	// Kita perlu cara untuk memberitahu pengguna bahwa server sudah berjalan
	// dan bagaimana cara menghentikannya.
	// Karena server sekarang berjalan di Goroutine, fungsi ini tidak lagi memblokir.
	// Kita akan kembali ke menu utama, dan tambahkan opsi "Stop API Server" di main.go.
}

// StopProductAPIServer menghentikan HTTP server secara graceful.
// Fungsi ini akan dipanggil dari main.go ketika user memilih opsi stop.
func StopProductAPIServer() {
	if serverInstance != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Memberi waktu 5 detik untuk shutdown
		defer cancel()                                                          // Pastikan context dibatalkan

		if err := serverInstance.Shutdown(ctx); err != nil {
			log.Printf("LOG: Error shutting down API server: %v", err)
			fmt.Println("Error: Gagal menghentikan server API.")
		} else {
			fmt.Println("Server API berhasil dihentikan. Kembali ke menu utama.")
		}
		serverInstance = nil // Reset instance server
	} else {
		fmt.Println("Server API tidak sedang berjalan.")
	}
}

// IsProductAPIRunning memeriksa apakah server API sedang berjalan.
// Berguna untuk menampilkan/menyembunyikan opsi "Start/Stop" di menu.
func IsProductAPIRunning() bool {
	return serverInstance != nil
}
