package main

import (
	"fmt"

	wallet "github.com/ainizoda/go-lessons-wallet/pkg/wallet"
)

func main() {
	svc := wallet.Service{}
	account, err := svc.RegisterAccount("+992502030070")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 500)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(account.Balance)
	payment, err := svc.Pay(1, 64, "food")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(account.Balance, payment.Category)
	svc.Reject(payment.ID)
	fmt.Println(account.Balance, payment.Category)
}
