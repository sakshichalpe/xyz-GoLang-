package main

import f "fmt"

func main() { //{a:1,b:2,c:3,d:4,e:5}

	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"d": 4, "e": 5, "a": 1, "c": 3}

	f.Println(map1)
	f.Println(map2)

	for key, value := range map2 {
		map1[key] = value
	}
	f.Println("Map after Merge =>", map1)

}
