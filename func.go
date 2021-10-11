package main

import "fmt"

func display(n string) int {
	fmt.Println(n, "is studing at TQ")
	return len(n)
}

func main() {
	fmt.Println("We are in main")
	fmt.Println("")
	var name string
	fmt.Scan(&name)
	fmt.Println("")
	var length = display(name)
	fmt.Println(length)
	fmt.Println("Back to main")
}
