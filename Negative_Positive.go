package main

import "fmt"

func main() {
	var size, i int

	fmt.Print("enter the positive Negative array size = ")
	fmt.Scan(&size)

	posNegArr := make([]int, size)

	fmt.Print("enter positive Negative array elements  = ")
	for i = 0; i < size; i++ {
		fmt.Scan(&posNegArr[i])
	}

	posCount := 0
	negCount := 0

	for i = 0; i < size; i++ {
		if posNegArr[i] >= 0 {
			posCount++
		} else {
			negCount++
		}
	}
	fmt.Println("Total Number of Positive No : ", posCount)
	fmt.Println("Total Number of Negative No : ", negCount)
}
