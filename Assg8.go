package main

import "fmt"

func main() {

	map_1 := map[string]int{"A": 100, "B": 100, "C": 300, "D": 400, "E": 300, "F": 300}
	fmt.Println(map_1)
	value, ok := map_1["104"]
	if ok {
		fmt.Println("Key found value is: ", value)
	} else {
		fmt.Println("Key not found")
		map_1["Mr.L"] = 104
	}
	fmt.Println(map_1)
}
