// mini-projects/calculator-app/calculator-app.go
package calculator_app // Mendeklarasikan package ini sebagai 'calculator_app'

import (
	"bufio"   // Untuk membaca input dari pengguna (misalnya dari os.Stdin)
	"fmt"     // Untuk mencetak output ke terminal (misalnya hasil perhitungan)
	"math"    // Untuk operasi matematika seperti pangkat (math.Pow)
	"strconv" // Untuk konversi string ke tipe numerik (misalnya float64)
	"strings" // Untuk manipulasi string (misalnya, TrimSpace, Fields)
)

// history adalah slice (array dinamis) yang menyimpan riwayat hasil perhitungan.
// Variabel ini bersifat internal untuk package calculator_app (huruf awal kecil).
var history []string

// hitung adalah fungsi internal untuk package ini yang melakukan operasi matematika dasar.
// Fungsi ini tidak diekspor (huruf awal kecil) karena hanya digunakan di dalam package ini.
func hitung(a float64, op string, b float64) (float64, error) {
	switch op {
	case "+":
		return a + b, nil // Penjumlahan
	case "-":
		return a - b, nil // Pengurangan
	case "*":
		return a * b, nil // Perkalian
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("pembagian dengan nol tidak valid") // Menangani pembagian dengan nol
		}
		return a / b, nil // Pembagian
	case "^":
		return math.Pow(a, b), nil // Pangkat (menggunakan math.Pow)
	default:
		return 0, fmt.Errorf("operator tidak dikenali") // Operator yang tidak didukung
	}
}

// RunCalculatorCLI adalah fungsi utama yang menjalankan aplikasi kalkulator CLI.
// Fungsi ini diekspor (huruf awal kapital 'R') sehingga bisa dipanggil dari package 'main'.
// Parameter 'reader' diperlukan untuk membaca input dari pengguna,
// agar fungsi ini bisa diintegrasikan dengan sistem input di main.go.
func RunCalculatorCLI(reader *bufio.Reader) {
	// Informasi awal ke pengguna
	fmt.Println("\n== Kalkulator CLI Sederhana ==")
	fmt.Println("Ketik 'exit' untuk keluar, 'history' untuk melihat riwayat.")
	fmt.Println("Operator yang didukung: +  -  * /  ^")

	for { // Loop utama kalkulator, akan terus berjalan sampai pengguna mengetik 'exit'
		// Meminta pengguna memasukkan ekspresi
		fmt.Print("\nMasukkan ekspresi (contoh: 5 + 3 - 1): ")
		input, _ := reader.ReadString('\n') // Membaca input sampai karakter newline
		input = strings.TrimSpace(input)    // Menghapus spasi di awal/akhir dan newline

		// Cek jika pengguna ingin keluar dari kalkulator
		if input == "exit" {
			fmt.Println("Keluar dari Kalkulator. Sampai jumpa!")
			break // Keluar dari loop for, menghentikan fungsi RunCalculatorCLI
		}

		// Cek jika pengguna ingin melihat riwayat perhitungan
		if input == "history" {
			if len(history) == 0 { // Jika riwayat kosong
				fmt.Println("Belum ada riwayat.")
			} else { // Jika ada riwayat
				fmt.Println("Riwayat Perhitungan:")
				for i, h := range history { // Iterasi dan tampilkan setiap riwayat
					fmt.Printf("%d. %s\n", i+1, h)
				}
			}
			continue // Lanjut ke iterasi berikutnya (meminta input lagi)
		}

		// Memecah input menjadi bagian-bagian (token) berdasarkan spasi
		tokens := strings.Fields(input)
		// Validasi dasar: jumlah token harus ganjil (angka operator angka operator angka...)
		if len(tokens)%2 == 0 {
			fmt.Println("Format ekspresi salah. Harus: angka operator angka ...")
			continue // Lanjut ke iterasi berikutnya
		}

		// Parsing angka pertama dari ekspresi
		hasil, err := strconv.ParseFloat(tokens[0], 64) // Konversi string ke float64
		if err != nil {
			fmt.Println("Angka pertama tidak valid:", tokens[0])
			continue // Lanjut ke iterasi berikutnya
		}

		// Lakukan perhitungan dari kiri ke kanan untuk sisa ekspresi
		for i := 1; i < len(tokens); i += 2 { // Loop melalui operator dan angka berikutnya
			op := tokens[i]                                // Ambil operator (misalnya "+", "-")
			angkaStr := tokens[i+1]                        // Ambil string angka berikutnya
			angka, err := strconv.ParseFloat(angkaStr, 64) // Konversi string angka ke float64
			if err != nil {
				fmt.Println("Angka tidak valid:", angkaStr)
				// Jika ada error parsing angka, hentikan perhitungan untuk ekspresi ini
				break
			}

			// Panggil fungsi hitung untuk melakukan operasi
			hasil, err = hitung(hasil, op, angka)
			if err != nil {
				fmt.Println("Error perhitungan:", err)
				// Jika ada error perhitungan (misal pembagian nol), hentikan
				break
			}
		}

		// Cek apakah ada error yang membuat loop perhitungan terhenti
		if err != nil {
			continue // Jika ada error di tengah perhitungan, lompat ke input berikutnya
		}

		// Format hasil output: tanpa desimal jika bilangan bulat, dua desimal jika pecahan
		var output string
		if hasil == float64(int(hasil)) { // Cek apakah hasil adalah bilangan bulat
			output = fmt.Sprintf("%s = %d", input, int(hasil)) // Format sebagai integer
		} else {
			output = fmt.Sprintf("%s = %.2f", input, hasil) // Format dengan dua desimal
		}

		fmt.Println("Hasil:", output)     // Tampilkan hasil perhitungan
		history = append(history, output) // Tambahkan hasil ke riwayat
	}
}
