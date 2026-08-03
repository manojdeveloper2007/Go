package main

import "fmt"

//variadic function means we can pass infinity arguments in functions

func average(nums ...int) {
	var avg float64
	var sum float64 = 0

	for _, n := range nums {
		sum += float64(n)
	}
	avg = sum / float64(len(nums))
	fmt.Println("length : ", len(nums))
	fmt.Println("Sum : ", sum)
	fmt.Println("Average : ", avg)
}

func main() {
	avg := average

	avg(1, 7, 9, 2, 5, 6, 12, 4)
}
