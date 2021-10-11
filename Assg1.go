package main

import (
	f "fmt"
)

func main() {
	sliceOne := []int{1, 2, 3}
	sliceTwo := []int{8, 9}

	result := append(sliceOne, sliceTwo...)
	f.Printf("After concatenate: %v\n", result)

}
