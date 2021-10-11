//Number of month to Days in month

package main

import "fmt"

func main() {
	var monthNo = 1
	var year int
	fmt.Println("Enter Month No :")
	fmt.Scan(&monthNo)

	switch monthNo {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 Days in Month")
	case 4, 6, 9, 11:
		fmt.Println("30 Days in Month")
	case 2:
		fmt.Println("Enter Year :")
		fmt.Scan(&year)
		if (year%4 == 0) && ((year%400 == 0) || (year%100 != 0)) {
			fmt.Println("Its a leap year 29 Days in a Month")
		} else {
			fmt.Println("Its not a leap year 28 Days in a Month")
		}
	}
}
