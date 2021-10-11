package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}                             //1 slice
	s2 := make([]int, len(s1))                              //2 slice
	fmt.Println("Original S1:",    s1, " Orignal S2", s2)
	fmt.Println()
	copy(s2, s1)                                           //copy method
	fmt.Println("After copy s1 in s2")
	fmt.Println("S1:", s1, "S2:", s2)                      // s1->[1 2 3 4 5]  s2->[1 2 3 4 5]

	fmt.Println("After making change in s2 : ")
	s2[1] = 10                                              // changing s2 does not affect s1
	fmt.Println("S1", s1, "S2", s2)                        // s1->[1 2 3 4 5]  s2->[1 10 3 4 5]
	fmt.Println("S1 is not affected")
}
