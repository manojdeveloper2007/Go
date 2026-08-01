package main

import "fmt"

func main() {
	arr := []int{12, 2, 1, 6, 4}

	sum := 0

	for _, num := range arr {
		sum += num
	}
	fmt.Println("Sum of the Slices is ", sum)
}
