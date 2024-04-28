package main

import wallet "github.com/ainizoda/go-lessons-wallet/pkg/wallet"

func main() {
	svc := wallet.Service{}
	svc.RegisterAccount("+992502030070")
}
