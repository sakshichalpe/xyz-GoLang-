package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Enter a number")
	fmt.Scan(&num)
	if num%5 == 0 && num%11 == 0 {
		fmt.Print("no is divisible by 5 nd 11")
	} else {
		fmt.Println("no is not divisible by 5 nd 11")
	}
}
