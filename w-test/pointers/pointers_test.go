package pointers

import (
	"fmt"
	"testing"
)

type Bitcoin int

// Implement Stringer interface
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

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

	fmt.Printf("got %s want %s\n", got, want)
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
