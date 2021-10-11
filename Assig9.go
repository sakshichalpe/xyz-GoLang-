package main

import f "fmt"

func main() {
	var Map_1 = map[int]string{101: "Mr.X", 102: "Mr.Y", 103: "Mr.Z", 104: "Mr.X", 105: "Mr.Z"}

	f.Println("Original Map", Map_1)

	//1 way
	for k := range Map_1 {
		delete(Map_1, k)
	}
	f.Println("After Deleting", Map_1)

	//2 way
	Map_1 = make(map[int]string)
	f.Println("After deleting", Map_1)
}
