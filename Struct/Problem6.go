package main

import "fmt"

// anonymous struct
func main() {
	student := struct {
		name string
		age  int
	}{
		name: "manoj",
		age:  18,
	}

	fmt.Println(student.name, " : ", student.age)
}
