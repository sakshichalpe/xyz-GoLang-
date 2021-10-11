//Write a code to display last element from slice
package main

import "fmt"

func main() {

	// Declaring & Defining the Slice
	sliceOfInt := []int{2, 3, 4, 5, 6}

	// Printing the Slice
	fmt.Println("Slice: ", sliceOfInt)

	// Accessing zeroth index
	// i.e. first element
	// first := sliceOfInt[0]

	// // Printing first element
	// fmt.Printf("First element: %d\n", first)

	// Accessing length(slice)-1
	// index i.e. last
	// element
	last := sliceOfInt[len(sliceOfInt)-1]

	// Printing last element
	fmt.Printf("Last element: %v\n", last)

}
