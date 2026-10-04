package main

import "fmt"

func main() {
	var Gb, Gp, Bl, Pt, jam_lembur int
	// Masukan Gaji Pokok dan Jam lembur
	fmt.Scan(&Gp, &jam_lembur)

	Bl = 45000 * jam_lembur // Menghitung Bonus lembur
	Pt = ((2 * Gp) / 100) + ((35 * Gp) / 1000) //  Menghitung Potongan

	Gb = Gp + Bl - Pt // Hasil Gaji pokok dengan cara mennghitung Gp + Bl - Pt
	// Menampilkan hasil perhitungan dari Gaji Pokok
	fmt.Println(Gb)
}