package main

import "fmt"

func main(){
	var kecepatan, menit, jam, total_jarak int

	// Masukan Kecepatan dalam bentuk bilangan bulat
	fmt.Scan(&kecepatan)
	//Proses
	total_jarak = 100 + 60 + 170 // Menghitung total jarak
	menit = (total_jarak * 60) / kecepatan // Menghitung dalam bentuk Menit
	jam = menit / 60 // Menghitung dalam bentuk jam
	menit = menit % 60 // Menghitung sisa menit
	// Hasil Output
	fmt.Println(jam, menit)
}