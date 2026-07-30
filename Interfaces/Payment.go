package main

import "fmt"

// creating interface
type Payment interface {
	payment(amount int)
}

// upi
type Upi struct {
	amount int
}

// atm
type Atm struct {
	amount int
}

// credit card
type CreditCard struct {
	amount int
}

func (u Upi) payment(amount int) {
	fmt.Println("Amount Pay Succesfully ", amount)
}

func (a Atm) payment(amount int) {
	fmt.Println("Amount Pay Succesfully ", amount)
}

func (c CreditCard) payment(amount int) {
	fmt.Println("Amount Pay Succesfully ", amount)
}

func main() {
	var p Payment

	upi := Upi{1200}
	atm := Atm{2000}
	creditcard := CreditCard{500}

	p = upi
	p.payment(upi.amount)

	p = atm
	p.payment(atm.amount)

	p = creditcard
	p.payment(creditcard.amount)
}
