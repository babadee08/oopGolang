package account

import "fmt"

type Account struct {
}

func (a Account) AvailableFunds() float32 {
	fmt.Println("Listing available funds")
	return 0
}
func (a Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment")
	return true
}

type Acc struct {
}

func (a Acc) AvailableFunds() float32 {
	fmt.Println("Listing available funds")
	return 0
}

type CreditAccount struct {
	Account
}

type HybridAcc struct {
	Acc
	Account
}

func (h HybridAcc) AvailableFunds() float32 {
	return h.Account.AvailableFunds() + h.Acc.AvailableFunds()
}
