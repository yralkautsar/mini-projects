// mini-projects/downloader-app/downloader-app.go
package parallel_downloader_app // Mendeklarasikan package ini sebagai 'parallel_downloader_app'

import (
	"bufio"    // Untuk membaca input dari pengguna (misalnya URL, nama file)
	"fmt"      // Untuk fungsi input/output seperti Println
	"io"       // Untuk operasi input/output (misalnya membaca dan menulis data stream)
	"net/http" // Untuk melakukan request HTTP ke server
	"os"       // Untuk berinteraksi dengan sistem operasi (misalnya membuat/menulis file, menghapus file)
	"strconv"  // Untuk konversi string ke angka dan sebaliknya
	"strings"  // Untuk manipulasi string (misalnya, menghapus spasi/newline)
	"sync"     // Untuk WaitGroup, agar kita bisa menunggu Goroutine selesai
	// Package 'time' tidak diperlukan di sini karena tidak ada time.Sleep() yang spesifik di logika inti downloadPart.
	// Jika ada, akan diimport.
)

// downloadPart adalah fungsi yang akan dijalankan oleh setiap Goroutine
// untuk mengunduh sebagian kecil dari file.
// 'url': URL file yang akan diunduh.
// 'partNum': Nomor bagian (misalnya, bagian ke-0, ke-1, dst.) untuk identifikasi.
// 'startByte': Byte awal dari bagian yang akan diunduh (termasuk).
// 'endByte': Byte akhir dari bagian yang akan diunduh (termasuk).
// 'outputFile': Nama file sementara tempat bagian ini akan disimpan di disk.
// 'wg': Pointer ke WaitGroup untuk memberi tahu Goroutine utama ketika selesai.
// 'errorCh': Channel untuk melaporkan error kembali ke Goroutine utama.
func downloadPart(url string, partNum int, startByte, endByte int64, outputFile string, wg *sync.WaitGroup, errorCh chan error) {
	defer wg.Done() // Pastikan wg.Done() dipanggil ketika Goroutine ini selesai, baik sukses atau error.

	fmt.Printf("[Bagian %d] Memulai download dari byte %d sampai %d...\n", partNum, startByte, endByte)

	// Membuat HTTP Request baru dengan metode GET.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Menggunakan %v untuk menampilkan error, menghindari masalah %w jika tidak ada error yang dibungkus.
		errorCh <- fmt.Errorf("gagal membuat request HTTP untuk bagian %d: %v", partNum, err)
		return
	}

	// Menambahkan header "Range" untuk meminta sebagian file saja dari server.
	// Contoh format: Range: bytes=0-999 (untuk 1000 byte pertama)
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", startByte, endByte))

	// Melakukan request HTTP menggunakan HTTP client default.
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		errorCh <- fmt.Errorf("gagal melakukan request HTTP untuk bagian %d: %v", partNum, err)
		return
	}
	defer resp.Body.Close() // Pastikan body response ditutup setelah selesai membaca.

	// Memeriksa status kode HTTP dari respons server.
	// Kode 206 (Partial Content) berarti server berhasil mengirim sebagian file.
	// Kode 200 (OK) bisa terjadi jika server tidak mendukung Range Request dan mengirim seluruh file.
	if resp.StatusCode != http.StatusPartialContent && resp.StatusCode != http.StatusOK {
		errorCh <- fmt.Errorf("server mengembalikan status %d untuk bagian %d", resp.StatusCode, partNum)
		return
	}

	// Membuat file sementara di disk untuk menyimpan bagian yang diunduh ini.
	file, err := os.Create(outputFile)
	if err != nil {
		errorCh <- fmt.Errorf("gagal membuat file %s untuk bagian %d: %v", outputFile, partNum, err)
		return
	}
	defer file.Close() // Pastikan file ditutup setelah selesai menulis.

	// Menyalin data yang diunduh dari body response HTTP ke file sementara.
	bytesWritten, err := io.Copy(file, resp.Body)
	if err != nil {
		errorCh <- fmt.Errorf("gagal menulis data ke file %s untuk bagian %d: %v", outputFile, partNum, err)
		return
	}

	fmt.Printf("[Bagian %d] Download selesai. Ukuran: %d bytes. Disimpan di: %s\n", partNum, bytesWritten, outputFile)
}

// RunParallelDownloaderCLI adalah fungsi utama yang menjalankan aplikasi Parallel File Downloader CLI.
// Fungsi ini diekspor (huruf awal kapital 'R') sehingga bisa dipanggil dari package 'main'.
// Parameter 'reader' diperlukan untuk membaca input dari pengguna (URL, nama file, jumlah bagian).
func RunParallelDownloaderCLI(reader *bufio.Reader) {
	fmt.Println("\n--- Go Parallel File Downloader ---")

	// --- Konfigurasi Awal (diambil dari input pengguna) ---
	fmt.Print("Masukkan URL file (contoh: https://speed.cloudflare.com/__down?bytes=10000000): ")
	fileURL, _ := reader.ReadString('\n')
	fileURL = strings.TrimSpace(fileURL)

	fmt.Print("Masukkan nama file output (contoh: downloaded_10MB.bin): ")
	outputFileName, _ := reader.ReadString('\n')
	outputFileName = strings.TrimSpace(outputFileName)

	fmt.Print("Masukkan jumlah bagian paralel (contoh: 4): ")
	numPartsStr, _ := reader.ReadString('\n')
	numPartsStr = strings.TrimSpace(numPartsStr)
	numParts, err := strconv.Atoi(numPartsStr)
	if err != nil || numParts <= 0 {
		fmt.Println("Jumlah bagian tidak valid atau kosong. Menggunakan default 4.")
		numParts = 4
	}

	fmt.Printf("Mencoba mengunduh file dari: %s\n", fileURL)
	fmt.Printf("Menggunakan %d Goroutine paralel.\n", numParts)

	// --- Step 1: Mendapatkan Ukuran File Total (Metadata) ---
	// Melakukan HEAD request untuk mendapatkan header file tanpa mengunduh seluruh body.
	resp, err := http.Head(fileURL) // http.Head() hanya mengambil header, lebih cepat.
	if err != nil {
		fmt.Printf("Error: Gagal mendapatkan header file: %v\n", err)
		return // Keluar dari fungsi jika gagal mendapatkan header
	}
	defer resp.Body.Close() // Pastikan body response ditutup setelah selesai.

	// Memeriksa apakah server mengembalikan status OK (200).
	// Ini penting untuk memastikan file ditemukan sebelum mencoba mengunduh.
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Server mengembalikan status %d. File mungkin tidak ditemukan atau tidak dapat diakses.\n", resp.StatusCode)
		return // Keluar jika status bukan OK
	}

	// Mendapatkan ukuran file dari header "Content-Length".
	contentLengthStr := resp.Header.Get("Content-Length")
	if contentLengthStr == "" {
		fmt.Println("Error: Header Content-Length tidak ditemukan. Tidak bisa menentukan ukuran file.")
		return // Keluar jika ukuran file tidak diketahui
	}
	fileSize, err := strconv.ParseInt(contentLengthStr, 10, 64) // Konversi string ke integer 64-bit
	if err != nil {
		fmt.Printf("Error: Gagal mengkonversi Content-Length ke angka: %v\n", err)
		return // Keluar jika konversi gagal
	}
	fmt.Printf("Ukuran file total: %d bytes\n", fileSize)

	// --- Step 2: Menghitung Rentang Byte untuk Setiap Bagian ---
	partSize := fileSize / int64(numParts) // Ukuran ideal setiap bagian
	var wg sync.WaitGroup                  // WaitGroup untuk menunggu semua Goroutine pengunduh selesai
	// Channel buffered untuk mengumpulkan error dari Goroutine.
	// Buffer sebesar 'numParts' agar Goroutine tidak blocking saat mengirim error.
	errorCh := make(chan error, numParts)

	// Membuat slice (array dinamis) untuk menyimpan nama file sementara dari setiap bagian.
	tempFiles := make([]string, numParts)

	// Memulai Goroutine untuk setiap bagian file
	for i := 0; i < numParts; i++ {
		startByte := int64(i) * partSize
		endByte := startByte + partSize - 1 // Byte akhir bagian ini (inklusif)

		// Untuk bagian terakhir, pastikan mencakup sisa byte yang mungkin ada
		// agar tidak ada data yang terlewat.
		if i == numParts-1 {
			endByte = fileSize - 1
		}

		// Membuat nama file sementara untuk setiap bagian (misal: my_file.zip.part0)
		tempFileName := fmt.Sprintf("%s.part%d", outputFileName, i)
		tempFiles[i] = tempFileName // Simpan nama file sementara di slice

		wg.Add(1) // Menambahkan 1 ke WaitGroup untuk setiap Goroutine yang akan dibuat
		// Menjalankan fungsi downloadPart sebagai Goroutine.
		// Setiap Goroutine akan mengunduh bagiannya secara paralel.
		go downloadPart(fileURL, i, startByte, endByte, tempFileName, &wg, errorCh)
	}

	// Goroutine terpisah untuk menutup channel error.
	// Ini penting agar loop `for err := range errorCh` di bawah bisa berhenti
	// setelah semua Goroutine download selesai dan tidak ada error lagi yang akan dikirim.
	go func() {
		wg.Wait()      // Tunggu sampai semua Goroutine download memanggil wg.Done()
		close(errorCh) // Setelah semua selesai, tutup channel error
	}()

	// --- Step 3: Menunggu Semua Goroutine Selesai & Menangani Error ---
	// Menerima error dari Goroutine melalui channel.
	// Loop ini akan berjalan sampai channel 'errorCh' ditutup.
	var downloadErrors []error // Slice untuk menyimpan semua error yang terjadi
	for err := range errorCh { // Menerima error dari channel
		downloadErrors = append(downloadErrors, err) // Tambahkan error ke slice
	}

	// Jika ada error yang terkumpul, tampilkan dan keluar.
	if len(downloadErrors) > 0 {
		fmt.Println("\nError saat mengunduh bagian:")
		for _, err := range downloadErrors {
			fmt.Println("-", err) // Tampilkan setiap error
		}
		fmt.Println("Download gagal karena error pada satu atau lebih bagian.")
		// Hapus file-file sementara yang mungkin sudah terunduh sebagian
		for _, f := range tempFiles {
			os.Remove(f) // Hapus file sementara
		}
		return // Keluar dari fungsi jika ada error
	}

	fmt.Println("\nSemua bagian berhasil diunduh. Memulai penggabungan...")

	// --- Step 4: Menggabungkan Bagian-bagian File ---
	// Membuat file akhir yang akan berisi gabungan semua bagian.
	finalFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Printf("Error: Gagal membuat file akhir %s: %v\n", outputFileName, err)
		return // Keluar jika gagal membuat file akhir
	}
	defer finalFile.Close() // Pastikan file akhir ditutup setelah selesai.

	// Menggabungkan setiap bagian file secara berurutan.
	for i := 0; i < numParts; i++ {
		partFileName := tempFiles[i]
		partFile, err := os.Open(partFileName) // Membuka file bagian sementara
		if err != nil {
			fmt.Printf("Error: Gagal membuka bagian %s: %v\n", partFileName, err)
			return // Keluar jika gagal membuka bagian
		}
		defer partFile.Close() // Pastikan file bagian ditutup

		// Menyalin isi dari file bagian ke file akhir.
		bytesCopied, err := io.Copy(finalFile, partFile)
		if err != nil {
			fmt.Printf("Error: Gagal menyalin data dari bagian %s ke file akhir: %v\n", partFileName, err)
			return // Keluar jika gagal menyalin data
		}
		fmt.Printf("Menggabungkan %s (%d bytes)...\n", partFileName, bytesCopied)
		os.Remove(partFileName) // Hapus file bagian setelah berhasil digabungkan
	}

	fmt.Printf("\n--- File '%s' berhasil diunduh dan digabungkan! ---\n", outputFileName)
	fmt.Printf("Total ukuran file: %d bytes\n", fileSize) // Tampilkan ukuran total file yang diunduh
}
