package main

import "fmt"

func arr_int(arr1 [5]int) [5]int { //0
	return arr1
}
func arr_float64(arr2 [5]float64) [5]float64 { //0
	return arr2
}
func arr_float32(arr3 [5]float32) [5]float32 { //0
	return arr3
}

func main() {
	var arr1 [5]int
	var arr2 [5]float64
	var arr3 [5]float32

	fmt.Println(arr_int(arr1))
	fmt.Println(arr_float64(arr2))
	fmt.Println(arr_float32(arr3))
}
