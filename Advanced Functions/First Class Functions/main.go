package main

import "fmt"

//first class functions means we can store function as a variable then use

func greet() {
	fmt.Println("Hello , Buddy !")
}

func sayHello() {
	fmt.Println("hello")
}

func sayBye() {
	fmt.Println("bye")
}

func main() {
	//	we can store function in variable
	//g := greet
	//g()

	//	swap function using variable

	hello := sayHello
	bye := sayBye

	temp := hello

	hello = bye

	bye = temp

	hello()
	bye()

}
