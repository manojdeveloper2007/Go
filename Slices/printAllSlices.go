package main

import "fmt"

func main() {
	//	declare using var

	var arr []int = []int{1, 2, 3, 4, 5}

	//	print using for loop

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//	Print using range
	for index, value := range arr {
		fmt.Println(index, " : ", value)
	}
}
