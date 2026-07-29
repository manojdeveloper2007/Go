package main

import "fmt"

type Bank struct {
	holder  string
	balance int
}

func deposit(bank *Bank, amount int) {
	if bank.balance < amount {
		fmt.Println("insufficient balance")
	} else {
		bank.balance -= amount
		fmt.Printf("Amount %drs deposit successfully from %s holder\n", amount, bank.holder)
	}
}

func credit(bank *Bank, amount int) {
	bank.balance += amount
	fmt.Printf("Amount credited Successfully to your account\n")
}

func main() {

	holder1 := Bank{
		holder:  "Manoj",
		balance: 100,
	}

	fmt.Println(holder1.balance)
	deposit(&holder1, 50)
	fmt.Println(holder1.balance)
	credit(&holder1, 500)
	fmt.Println(holder1.balance)
	deposit(&holder1, 550)
	fmt.Println(holder1.balance)
}
