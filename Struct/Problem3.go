package main

import "fmt"

type Car struct {
	brand string
	model string
	price int
}

func showDetails(car Car) {
	fmt.Printf("Car Brand : %s\nModel : %s\nPrice : %d\n", car.brand, car.model, car.model)
}

var car = Car{
	brand: "BMW",
	model: "bmw M5",
	price: 14000000,
}

func main() {
	showDetails(car)
}
