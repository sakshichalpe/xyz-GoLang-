package main

import f "fmt"

func main() {
	slice := make([]string, 4, 5)

	//add elemets in slice
	slice[0] = "Prakash"
	slice[1] = "Chanda"
	slice[2] = "Sanket"
	slice[3] = "Sakshi"

	//Iterate Slic using For Loop
	f.Println("--------Through For loop--------")
	f.Println("-------slice position nd element--------")
	for i := 0; i < len(slice); i++ {
		f.Println(i, "=", slice[i])
	}
	//Iterate Slic using Range
	f.Println("--------Through Range loop--------")
	f.Println("-------slice position nd element--------")
	for index, value := range slice {
		f.Println(index, "=", value)
	}
}
