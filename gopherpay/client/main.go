package main

import (
	"fmt"
	"oopGolang/gopherpay/account"
	"oopGolang/gopherpay/payment"
)

type CredAcc struct{}

func (c *CredAcc) processPayment(amount float32) {
	fmt.Println("Processing Cred kard payment...")
}

func CreateCredAcc(chargeCh chan float32) *CredAcc {
	credAcc := &CredAcc{}
	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			credAcc.processPayment(amount)
		}
	}(chargeCh)
	return credAcc
}

func main() {

	credit := payment.CreateCreditAccount(
		"Debra Ingram",
		"5555-5555-5555-444",
		5,
		2021,
		345)

	fmt.Printf("Owner name: %v\n", credit.OwnerName())
	fmt.Printf("Card Number: %v\n", credit.CardNumber())

	err := credit.SetCardNumber("invalid")
	if err != nil {
		fmt.Printf("That did't work: %v\n", err)
	}

	fmt.Printf("Avialable Credit: %v\n", credit.AvailableCredit())

	var option payment.PaymentOption

	option = credit

	option.ProcessPayment(500)

	option = payment.CreateCashAccount()

	option.ProcessPayment(500)

	chargeCh := make(chan float32)
	CreateCredAcc(chargeCh)
	chargeCh <- 500

	ca := account.CreditAccount{}
	ca.AvailableFunds()
	ca.ProcessPayment(500)

	ha := account.HybridAcc{}
	ha.AvailableFunds()

	var a string
	_, err = fmt.Scanln(&a)
	if err != nil {
		return
	}
}
