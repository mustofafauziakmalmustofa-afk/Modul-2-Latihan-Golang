package main

import "fmt"

func main() {
	var N, hari int

	fmt.Print("Masukan jumlah hari hingga keberangkatan berupa bilangan bulat: ")
	fmt.Scan(&N)

	// Menghitung jumlah hari setelah setelah hari kamis
	hari = (4 + N - 1) % 7 + 1

	fmt.Print(hari)
}