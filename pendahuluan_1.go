package main

import "fmt"

func main() {
	var N, sisa, jam, menit, detik int
	
	// Masukan Jumlah Satuan Digit
	fmt.Print("Masukan jumlah detik: ")
	fmt.Scan(&N)

	// Proses Menghitung jam, menit dan detik
	jam = N / 3600
	sisa = N % 3600
	menit = sisa / 60
	detik = sisa % 60

	// Menampilkan Hasil Konversi Waktu
	fmt.Printf("N = %v: %v jam %v menit %v detik", N, jam, menit, detik)
}