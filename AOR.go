package main

import (
	"fmt"
)

func main() {
	var length, breadth float32
	fmt.Print("enter the length")
	fmt.Scan(&length)
	fmt.Print("enter the breadth")
	fmt.Scan(&breadth)

	var area = length * breadth

	fmt.Print("Area of the rectangle is: ", area)
}

//func rec() {
//	var length float32
//	var breadth float32
//	fmt.Print("enter the length")
//	fmt.Scan(&length)
//	fmt.Println("enter the breadth")
//	fmt.Scan(&breadth)
//
//	var area float32= length * breadth

//func main() {
//	rec()
//}
