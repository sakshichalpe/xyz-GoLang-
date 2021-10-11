package main

//Day of a week
import "fmt"

func main() {
	var day string
	fmt.Println("Enter a Day Name of week :")
	fmt.Scan(&day)

	switch day {
	case "Monday":
		fmt.Println("1st Day of a week")
	case "Tuesday":
		fmt.Println("2nd Day of a week")
	case "Wednesday":
		fmt.Println("3rd Day of a week")
	case "Thursday":
		fmt.Println("4th Day of a week")
	case "Friday":
		fmt.Println("5th Day of a week")
	case "Saturday":
		fmt.Println("6th Day of a week")
	case "Sunday":
		fmt.Println("7th Day of a week")
	}
}
