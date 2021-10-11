package max

import "fmt"

func max() {
	var a, b, c int
	fmt.Printf("Enter three number : %d %d %d")
	fmt.Scan(&a, &b, &c)
	//fmt.Print("")
	if a > b && a > c {
		fmt.Print("%d is greater", a)
	} else if b > a && b > c {
		fmt.Print("%d is greater", b)
	} else {
		fmt.Print("%d is greater", c)
	}
}
