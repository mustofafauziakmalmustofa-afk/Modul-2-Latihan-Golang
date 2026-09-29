package main

import "fmt"

func main() {
	var fahrenheit, celcius int

	fmt.Scan(&fahrenheit)

	celcius = (fahrenheit - 32) * 5/9

	fmt.Print(celcius)
}