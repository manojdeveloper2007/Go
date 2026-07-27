package main

import (
	"fmt"
)

func main() {
	// integers => int16,int32,int64,int8
	var num1 int = 97
	fmt.Printf("The Number 1 is %v and Type of the Number is %T\n", num1, num1)

	// unsigned integers
	var unsignNum uint = 100
	fmt.Printf("The Number Unsigned is %d and Type of the Number is %T\n", unsignNum, unsignNum)

	// boolean
	var isAlive bool = true
	fmt.Printf("He is still %t and Type of isAlive is %T\n", isAlive, isAlive)

	// format specifier => %v (handles every format automatically)
	// %s (strings) , %d (decimal numbers) , %f (floating) , %t (boolean)

	// String

	var name string = "Manoj"
	fmt.Printf("Hello %s ! How are you ? ", name)

	// without declare variable using types

	indianCaptain := "Mahendra Singh Dhoni"

	// we use Sprint() => which return string
	msg := fmt.Sprintf("This is our new Indian Captain %s", indianCaptain)

	fmt.Println(msg)

	// typecasting

	num2 := 3.14

	num3 := int(num2)

	fmt.Println("The number After type casting is ", num3)

	// format our floating numbers

	dec := 12.4520
	fmt.Printf("Only 2 digits After decimal : %.2f\n", dec)

	// constants
	const pi float64 = 3.14

	fmt.Println("constant : ", pi)

}
