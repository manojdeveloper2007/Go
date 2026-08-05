package main

import "fmt"

//anonymous function means these function doe not have any name

func main() {
	//	immediate call

	func() {
		fmt.Println("hello")
	}()

	f := func(a, b int) int {
		return a + b
	}

	fmt.Println(f(10, 2))

	res := func(a, b int) int {
		return a + b
	}(12, 5)

	fmt.Println(res)
}
