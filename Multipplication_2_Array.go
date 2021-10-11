package main

import "fmt"

func main() {
	var i int

	var multiArr1 [5]int
	var multiArr2 [5]int
	var MultiArr [5]int

	fmt.Print("Enter First Array Items ")
	for i = 0; i < 5; i++ {
		fmt.Scan(&multiArr1[i])
	}

	fmt.Print("Enter the Second Array Items  ")
	for i = 0; i < 5; i++ {
		fmt.Scan(&multiArr2[i])
	}

	fmt.Print("The Multip of Two Arrays = ")
	for i = 0; i < len(MultiArr); i++ {
		MultiArr[i] = multiArr1[i] * multiArr2[i]
		fmt.Print(MultiArr[i], "  ")
	}
	fmt.Println()
	
}
