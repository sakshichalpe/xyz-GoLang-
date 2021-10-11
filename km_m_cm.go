package main

import (
	"fmt"
)

func main() {
	var c, m, km int
	fmt.Print("enter number in Centimeter")
	fmt.Scan(&c)
	m = c / 100
	km = c / 100000
	fmt.Print("Converted in meter", m)
	fmt.Print("Converted in kelometer", km)
}
