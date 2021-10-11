package main

import (
	f "fmt"
	"sort"
)

// func createSlice(a [5]int) {
// 	var b []int = a[1:5]
// 	f.Println(b)
// }
func modifyslice(a [5]int) {
	f.Println("Before change", a)
	a[0] = 90
	f.Println("After Change", a)
}

func main() {
	var a [5]int = [5]int{81, 82, 83, 84, 85}
	//createSlice(a)
	modifyslice(a)
	var a1 []string = []string{"Sakshi"}
	a1 = append(a1, "Chalpe")
	f.Println(a1)
	f.Println(cap(a), len(a))

	//slice using literal
	var s1 = []int{65, 66, 67}
	f.Println("slice s1", s1)
	f.Println("slice s1 length", len(s1))
	f.Println("Slice s1 capacity", cap(s1))

	//sort slice
	//sort.Ints("Slice s1", s1)

	//slice using array
	var array = [5]string{"Vijay", "Kaushik", "Krishna", "Neha", "Raju"}
	s2 := array[2:]
	f.Println("slice s2", s2)
	f.Println("slice s1 length", len(s2))
	f.Println("Slice s1 capacity", cap(s2))
	sort.Strings(s2)
	//sort.Searchstrings(s2)

	//slice using Slice
	s3 := s2[1:3]
	f.Println("slice s2", s3)
	f.Println("slice s1 length", len(s3))
	f.Println("Slice s1 capacity", cap(s3))

	//creating empty slice
	s4 := []int{} //emplty slice
	if s4 == nil {
		f.Println("s4 is nil")
	}
	f.Println("slice s2", s4) //0
	f.Println("slice s1 length", len(s4))
	f.Println("Slice s1 capacity", cap(s4))

	p := new([5]int)[2:4]
	f.Println("p=", p)

	s44 := new([20]int)[5:15]
	f.Println("slice s2", s44) //0
	f.Println("slice s1 length", len(s44))
	f.Println("Slice s1 capacity", cap(s44))

	//cretae slic using make
	s5 := make([]int, 4)
	f.Println("slice s2", s5) //0
	f.Println("slice s1 length", len(s5))
	f.Println("Slice s1 capacity", cap(s5))

	//Iterate SLic using For Loop
	f.Println("Iterate SLic using For Loop")
	for i := 0; i < len(s2); i++ {
		f.Println(s2[i], i)
	}

	//Range for slice(position nd element)
	for index, value := range s2 {
		f.Println("%v is at index %d\n", value, index)
	}
	for _, value := range s2 {
		f.Println("%v  \n", value)
	}
	//change made in slice will make change in array also
	s2[0] = "vijay"
	f.Println(s2)
	f.Println(array)

}
