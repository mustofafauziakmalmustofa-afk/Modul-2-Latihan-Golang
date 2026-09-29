package main

import "fmt"

func main() {
	var ngueng int = 15
	alamatMemori := &ngueng

	fmt.Println(alamatMemori)
	fmt.Println(*alamatMemori)
}