package main

import (
	account "bank/accounts"
	"bank/customer"
	"fmt"
)

type CheckAccount interface {
	Withdraw(values ...float64) string
}

func BoletoPayment(account CheckAccount, value float64) {
	account.Withdraw(value)
}

func main() {
	userSilvia := customer.HolderUser{
		Name: "Silvia",
		CPF:  "01223456789",
		Job:  "Developer",
	}

	userGustavo := customer.HolderUser{
		Name: "Gustavo",
		CPF:  "12234567890",
		Job:  "Design",
	}

	silviaAccount := account.CheckingAccount{HolderUser: userSilvia}
	gustavoAccount := account.CheckingAccount{HolderUser: userGustavo}

	silviaAccount.Deposit(500)
	gustavoAccount.Deposit(600)

	fmt.Println(silviaAccount.Withdraw(100, 100, 100))
	fmt.Println(silviaAccount.Withdraw(100, -100))
	fmt.Println(silviaAccount.Withdraw(100, 100, 500))
	fmt.Println(silviaAccount.Funds())

	message, funds := silviaAccount.Deposit(100, 100)
	silviaAccount.Deposit(100, 100, -100)
	fmt.Println(message, funds)

	fmt.Println(gustavoAccount)
	fmt.Println(silviaAccount)

	status := silviaAccount.Transfer(100, &gustavoAccount)

	fmt.Println(status)
	fmt.Println(gustavoAccount)
	fmt.Println(silviaAccount)

	BoletoPayment(&gustavoAccount, 200)
	fmt.Println(gustavoAccount)
}
