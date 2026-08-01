package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 5, 0, 9, 3, 6}

	even, odd := 0, 0

	for _, num := range arr {
		if num%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	fmt.Println("Count of Even numbers in Slices is ", even)
	fmt.Println("Count of Odd numbers in Slices is ", odd)
}
