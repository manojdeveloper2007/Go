package main

import "fmt"

//higher order function means we can pass function as an argument and also it returns a function

func result(mark int) string {
	if mark >= 35 {
		return "Pass"
	}
	return "Fail"
}

func student(mark int, result func(mark int) string) string {
	return result(mark)
}

func main() {
	ans := student(33, result)
	fmt.Println(ans)
}
