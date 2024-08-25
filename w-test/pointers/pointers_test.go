package pointers

import (
	"fmt"
	"testing"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance is %p \n", &w.balance)

	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()

	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
