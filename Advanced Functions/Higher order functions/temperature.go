package main

import "fmt"

func temperature(celcius float64) func() float64 {
	return func() float64 {
		return (celcius * 1.8) + 32
	}
}

func main() {
	temp := temperature(35)
	fmt.Println("Fahrenheit : ", temp())
}
