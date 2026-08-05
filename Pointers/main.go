package main

import (
	"errors"
	"fmt"
)

type Bank struct {
	name    string
	balance int
}

func (b *Bank) credit(amount int) error {
	if amount < 0 {
		return errors.New("negative amount cannot be credit")
	}
	b.balance = b.balance + amount
	return nil
}

func (b *Bank) debit(amount int) {
	if b.balance < amount {
		fmt.Println("Insufficient balance")
	} else {
		b.balance -= amount
	}
}

func (b *Bank) getBalance() int {
	return b.balance
}

func main() {
	b := Bank{"Manoj", 200}

	fmt.Println(b.getBalance())
	b.debit(300)
	fmt.Println(b.getBalance())

	err := b.credit(500)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b.getBalance())
}
