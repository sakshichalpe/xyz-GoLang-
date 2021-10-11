//Calculator
package main

import "fmt"

func main() {
	fmt.Println("========Menu=======")
	fmt.Println("1. Addition")
	fmt.Println("2. Substraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	var a = 0.0
	var b = 0.0
	var choice = 0

	fmt.Println("Enter value of a :")
	fmt.Scan(&a)
	fmt.Println("Enter value of b :")
	fmt.Scan(&b)

	fmt.Println("Enter Your choice :")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Addition of", a, "and", b, "is", a+b)
	case 2:
		fmt.Println("substraction of", a, "and", b, "is", a-b)
	case 3:
		fmt.Println("Multiplication of", a, "and", b, "is", a*b)
	case 4:
		fmt.Println("Division of", a, "and", b, "is", a/b)
	}

}
