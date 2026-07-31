package main

import (
	"errors"
	"fmt"
)

// by creating a manual error
func divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("Cannot divide by Zero")
	}
	return num1 / num2, nil
}

func main() {
	var num1 float64 = 12
	var num2 float64 = 4

	ans, err := divide(num1, num2)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Answer : ", ans)

}
