package main

import "fmt"
import "math"

func main() {
	var r, luas float64

	fmt.Scan(&r)

	luas = math.Pi * r * r

	fmt.Printf("%.1f\n", luas)
}