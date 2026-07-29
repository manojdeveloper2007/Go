package main

import "fmt"

/*
Create a struct called Employee.

Fields:

id
name
salary

Create two employees and print them.Create a struct called Employee.

Fields:

id
name
salary

Create two employees and print them.
*/

type Employee struct {
	id     int
	name   string
	salary int64
}

func main() {
	employee1 := Employee{
		id:     8015,
		name:   "Manoj",
		salary: 200000,
	}

	employee2 := Employee{
		id:     8049,
		name:   "Sanjay",
		salary: 100000,
	}

	fmt.Printf("Hello %s,Welcome to our office\n", employee1.name)
	fmt.Printf("Hello %s,Welcome to our office\n", employee2.name)
}
