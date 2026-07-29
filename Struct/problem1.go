package main

import "fmt"

//Create a struct named Student with the following fields:
//
//name (string)
//age (int)
//mark (float64)
//
//Create one student and print all the values.

type Student struct {
	name string
	age  int
	mark float64
}

func main() {
	student1 := Student{
		name: "Manoj",
		age:  18,
		mark: 90.6,
	}
	fmt.Printf("Hello %s ,your age is %d and your mark is %f", student1.name, student1.age, student1.mark)
}
