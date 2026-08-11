package main

import "fmt"

type User struct {
	Name string
}

var user = User{"Manoj"}

// change username
func (u *User) changeName() {
	u.Name = "dhoni"
}

func main() {
	//	defer keyword which means it execute that line atlast
	defer fmt.Println(user.Name)
	defer fmt.Println("End of the Program")
	user.changeName()
	fmt.Println(user.Name)
	fmt.Println(user)
}
