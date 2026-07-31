package main

import (
	"errors"
	"fmt"
)

type Calculator interface {
	addition(num1, num2 int) int
	subtraction(num1, num2 int) int
	multiplication(num1, num2 int) int
	division(num1, num2 int) int
}

type operations struct {
	num1 int
	num2 int
}

func (operations) addition(num1, num2 int) int {
	return num1 + num2
}

func (operations) subtraction(num1, num2 int) int {
	return num1 - num2
}

func (operations) multiplication(num1, num2 int) int {
	return num1 * num2
}

func (operations) division(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("Cannot divide by Zero")
	}
	return num1 / num2, nil
}

func main() {
	var s = operations{12, 6}
	fmt.Printf("Addition of Two Numbers is %v\n", s.addition(s.num1, s.num2))
	fmt.Printf("Subtraction of Two Numbers is %v\n", s.subtraction(s.num1, s.num2))
	fmt.Printf("Multiplication of Two Numbers is %v\n", s.multiplication(s.num1, s.num2))

	div, err := s.division(s.num1, s.num2)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Division of two Numbers is %v\n", div)

}
