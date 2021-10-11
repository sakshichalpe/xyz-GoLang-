package main

import (
	"fmt"
)

func main() {
	var r float64
	const pi = 3.14
	fmt.Print("enter radius")
	fmt.Scan(&r)
	var area float64 = pi * r
	fmt.Print("Area of circle is: ", area)
	var circum float64 = 2 * pi * r * r
	fmt.Print("Circum of circle is: ", circum)
	var dia float64 = 2 * r
	fmt.Print("Diameter of circle is: ", dia)
}
