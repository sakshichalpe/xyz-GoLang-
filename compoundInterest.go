package main

import (
	"fmt"
	"math"
)

func main() {

	var Pamount, InterestRate, timePeriod, ciFuture, compoundI float64

	fmt.Print("Enter the Total or Principal Amount = ")
	fmt.Scanln(&Pamount)

	fmt.Print("Enter the Interest Rate = ")
	fmt.Scanln(&InterestRate)

	fmt.Print("Enter the Total number of Years = ")
	fmt.Scanln(&timePeriod)

	ciFuture = Pamount * (math.Pow((1 + InterestRate/100), timePeriod))
	compoundI = ciFuture - Pamount

	fmt.Println("\nThe Compound Interest  = ", compoundI)
	fmt.Println("The Future Compound Interest  = ", ciFuture)
}
