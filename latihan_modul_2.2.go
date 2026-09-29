package main

import "fmt"

func main() {
	var nama, kelas, nim string

	fmt.Scan(&nama, &kelas, &nim)

	fmt.Printf("Perkenalkan nama saya %v. Saya adalah seorang mahasiswa Telkom University Purwokerto, program studi Teknik Informatika dari kelas %v dengan NIM %v.", nama, kelas, nim)
}