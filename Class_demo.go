package main

import "fmt"

func oddeven(num int) bool {
	if num%2 == 0 {
		return true
		//fmt.Println("Even")
	} else {
		return false
		//	fmt.Println("Odd")
	}
}
func checkdigit() {
	//oddeven()
}
func createarray() {
	var num [3]int
	cities := [...]string{"pune", "Nashik"}

	num[0] = 35
	num[1] = 34
	fmt.Println(num)
	fmt.Println(cities)
	//metroCities := cities
}
func createArray1(b [5]int) {
	var n [5]int
	var sum = 0
	fmt.Println("Enter elements of array")
	for i := 0; i < len(n); i++ {
		fmt.Scan((&n[i]))
		sum += n[i]
	}
	for i := 0; i < len(n); i++ {
		fmt.Scan((n[i]))
	}
	fmt.Println(sum)
}
func sumofeven() {}

func main() {
	//var name string
	// var num int
	// fmt.Println("ENTER THE NUMBER")
	// fmt.Scan(&num)
	// var a = oddeven(num)
	// fmt.Println(a)
	//createArray1()
	//
	createSlice()
}

func displayArray(a [5]string) {
	fmt.Println("Array elements are")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
func createMultiDArray() {
	a := [3][5]int{{3, 4, 11, 7, 45},
		{4, 5, 2}, {7, 5, 8}}
	fmt.Println(a)
}

// func compare() {
// 	arr1 := [3]int{4, 5, 6}
// 	arr2 := [3]int{7, 8, 9}
// 	// arr1==arr2

func createSlice() {
	num := [...]int{34, 55, 3, 45, 2, 7, 2, 9}
	s1 := num[1:7]

	fmt.Println("slice contains", s1)
	fmt.Println("length of slice contains", len(s1))
	fmt.Println("Capacity of sloce contains", cap(s1))

}
