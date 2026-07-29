package main

import "fmt"

/*
	find the area of the circle using struct and display the area using function
*/

type Circle struct {
	radius float64
}

func area(c Circle) float64 {
	return 3.14 * (c.radius) * (c.radius)
}

func main() {
	c := Circle{
		radius: 2,
	}

	fmt.Printf("the radius of the circle is %f\n", area(c))
}
