package main

import "fmt"

// Pass by Value
func swap1(num1, num2 int) {
	num1, num2 = num2, num1
	fmt.Println("In Function")
	fmt.Println("Num1 : ", num1, "\nNum2 : ", num2, "\n")
}

// Pass by Reference
func swap2(num1 *int, num2 *int) {
	*num1, *num2 = *num2, *num1
	fmt.Println("In Function")
	fmt.Println("Num1 : ", *num1, "\nNum2 : ", *num2, "\n")
}

func main() {
	num1 := 10
	num2 := 20

	fmt.Printf("Num1 : %d \nNum2 : %d\n", num1, num2)
	swap2(&num1, &num2)
	fmt.Printf("Num1 : %d \nNum2 : %d\n", num1, num2)
}
