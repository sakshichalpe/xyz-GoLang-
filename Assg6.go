package main

import "fmt"

func main() {
	var map_1 = map[int]string{101: "Mr.X", 102: "Mr.Y", 103: "Mr.Z", 104: "Mr.X", 105: "Mr.Z"}
	fmt.Println(map_1)
	frequency := make(map[string]int)
	for _, teacher := range map_1 {
		if frequency[teacher] == 0 {
			frequency[teacher] = 1
		} else {
			frequency[teacher]++
		}
	}
	for teacher, count := range frequency {
		//fmt.Println(teacher, count)
		fmt.Print(teacher, ":", count, "  ") //(teacher, count)
	}
}
