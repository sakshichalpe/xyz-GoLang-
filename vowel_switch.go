package main

import "fmt"

func main() {

	var alphabet = "null"
	fmt.Print("Enter an alphabet :")
	fmt.Scan(&alphabet)

	switch alphabet {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		fmt.Println(alphabet, " is an vowel")
	default:
		fmt.Println(alphabet, " is a Consonant")
	}
}
