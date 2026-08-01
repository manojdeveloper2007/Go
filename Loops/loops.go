package main

import "fmt"

func main() {
	/*
		print
		1
		12
		123
		1234
		12345
	*/

	//In golang ,there is only for loop

	//using for loop
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}

	//	using while loop type
	i := 1

	for i <= 5 {
		fmt.Println(i)
		i++
	}
}
