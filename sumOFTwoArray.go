package main

import f "fmt"

func sumOfTwoArray(arr1 [5]int, arr2 [5]int) [5]int {
	var arr3 [5]int
	for i := 0; i < 5; i++ {
		arr3[i] = arr1[i] + arr2[i]
	}
	return arr3
}
func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{6, 7, 8, 9, 10}
	f.Println(sumOfTwoArray(arr1, arr2))
}
