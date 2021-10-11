package main

import (
	f "fmt"
)

func main() {
	f.Println("Enter choice:")
	f.Println("1: number prime or not")
	f.Println("2: number even or not")

	var s1 = []int{1, 2, 3, 4, 5}
	even := 0
	for _, num := range s1 {
		if num%2 == 0 {
			f.Print(num, "\t")
			even = even + num
		}
	}
}

// for i := 1; i < len(s1); i++ {
// 	if i%2 == 0 {
// 		f.Print(i, "\t")
// 		even = even + i
// 	}
// }
// f.Println("Even Numbers", even)
