package main

import f "fmt"

func sumofArray(arr [5]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + i
	}
	return sum
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	f.Println(sumofArray(arr))
}
