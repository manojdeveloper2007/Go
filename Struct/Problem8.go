package main

import "fmt"

type Person struct {
	name string
}

func (p Person) greet() string {
	msg := fmt.Sprintf("Hello , %s", p.name)
	return msg
}

func main() {
	p := Person{
		name: "Manoj",
	}
	fmt.Println(p.greet())
}
