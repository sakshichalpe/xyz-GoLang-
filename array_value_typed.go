package main

import f "fmt"

func array_value_typed() {
	a1 := [3]string{"pune", "Mumbai", "Banglore"}
	b1 := a1
	b1[1] = "Shimla"
	f.Println("1st array", a1)
	f.Println("2st array", b1)

}
func main() {
	array_value_typed()
}
