package main

import "fmt"

// for adding
func add(num1, num2 int) (sum int) {
	sum = num1 + num2
	return sum
}

// for subtracting
func sub(num1, num2 int) int {
	var sub int
	sub = num1 - num2
	return sub
}

func main() {
	var num1 int = 12
	var num2 int = 4
	var ans int
	ans = add(num1, num2)
	fmt.Println(ans)
	fmt.Println(sub(num1, num2))
}
