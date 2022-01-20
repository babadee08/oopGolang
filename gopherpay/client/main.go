package main

import (
	"fmt"
	"oopGolang/gopherpay/payment"
)

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
}
