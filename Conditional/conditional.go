package main

import "fmt"

func main() {
	var age int = 18

	if age >= 18 {
		fmt.Println("You are Eligible to vote")
	} else if age == 20 {
		fmt.Println("Hey Buddy")
	} else {
		fmt.Println("You are not eligible to vote")
	}

	//	using switch

	var day int = 2

	switch day {
	case 1:
		fmt.Println("Monday")

	case 2:
		fmt.Println("Tuesday")

	case 3:
		fmt.Println("Wednesday")

	default:
		fmt.Println("Other day")

	}

}
