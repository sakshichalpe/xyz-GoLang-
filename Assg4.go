package main

import (
	f "fmt"
)

func main() {
	var Slice1 = []string{"Ramesh", "Ganesh", "Arti", "Shivam"}
	f.Println("Original Slice is: ", Slice1)

	var index int = 0
	f.Printf("Enter a index number ")
	f.Scan(&index)

	Slice1 = append(Slice1[:index], Slice1[index+1:]...)
	f.Println("After deleting ", index+1, "element result is", Slice1)

	// Slice1 = Slice1[:len(Slice1)-1] //delete last
	// f.Println("After deleting Last element: ", Slice1)

	// Slice1 = append(Slice1[:2], Slice1[6:]...)
	// f.Println("Removed elements from 2-5    ", Slice1)
}
