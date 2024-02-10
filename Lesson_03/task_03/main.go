package main

import "fmt"

type BankAccount struct {
	account string
	owner   string
	balance int
}

func (b *BankAccount) Deposit(i int) {
	b.balance += i
}

func (b *BankAccount) Withdraw(o int) {

	if o > b.balance {
		fmt.Println("Недостаточно средств.")
	} else {
		b.balance -= o
	}

}

func (b *BankAccount) Balance() float64 {
	return float64(b.balance)
}

func main() {
	account := &BankAccount{
		owner:   "Иван",
		balance: 1000,
	}

	account.Deposit(500)
	fmt.Printf("%s: Баланс - $%.2f\n", account.owner, account.Balance())

	account.Withdraw(200)
	fmt.Printf("%s: Баланс - $%.2f\n", account.owner, account.Balance())

	account.Withdraw(1500)
	fmt.Printf("%s: Баланс - $%.2f\n", account.owner, account.Balance())
}
