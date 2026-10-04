package main

import "fmt"

func main() {
	var jumlah_koin, sisa, emas, perak, tembaga int

	// Masukan Koin
	fmt.Scan(&jumlah_koin)

	// Hitung Koin dalam bentuk emas, perak dan tembaga
	emas = jumlah_koin / 9
	sisa = jumlah_koin % 9
	perak = sisa / 3
	tembaga = jumlah_koin % 3
	// Menampilkan Hasil dari Perhitungan
	fmt.Println(emas, perak, tembaga)
}