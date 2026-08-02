package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	mp := make(map[string]student)
	mp["Manoj"] = student{
		name: "Manoj",
		age:  19,
	}

	mp["Suriya"] = student{
		name: "Suriya",
		age:  70,
	}

	for k := range mp {
		fmt.Println("Name : ", mp[k].name)
		fmt.Println("Age : ", mp[k].age)
	}
}
