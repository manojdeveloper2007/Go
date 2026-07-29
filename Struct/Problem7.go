package main

import "fmt"

//Nested structs

type Address struct {
	city  string
	state string
}

type Employeee struct {
	name    string
	age     int
	address Address
}

func main() {
	employee := Employeee{
		name: "Manoj",
		age:  18,
		address: Address{
			city:  "pudukkottai",
			state: "Tamilnadu",
		},
	}

	fmt.Println(employee.name, " : ", employee.age)
	fmt.Println(employee.address.city)
}
