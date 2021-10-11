package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Array:", arr)

	myslice := arr[5:] // Creating a slice

	fmt.Println("Last Five elemets are:", myslice) // Display slice

	// Display length & capacity of the slice
	fmt.Println("Length, Capacity of the slice is:", len(myslice), ",", cap(myslice), "respectively")

}
