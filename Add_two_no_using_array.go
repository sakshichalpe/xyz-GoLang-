package main

import f "fmt"

func add() { //
	var num1 int
	var num2 int
	f.Println("enter 1st the number ")
	f.Scan(&num1)
	f.Println("enter 2nd the number ")
	f.Scan(&num2)
	add := num1 + num2
	f.Println("Addition is:", add)
}

func test(x int) {
	f.Println(x)
}

func add1(x, y int) int {
	return x + y
}

func calc(x int, y int) (out1, out2 int) {
	out1 = x + y
	out2 = x - y
	return out1, out2
}

func main() {
	ans := add1(14, 7) //task 4
	f.Println(ans)

	test(5) //task 3

	num1 := 5
	num2 := 4
	result1, result2 := calc(num1, num2)
	f.Println(result1, result2) //task2

	add() //task 1
}
