package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	// operates on a pointer to a wallet, see Deposit definition
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	// %p prints memory address in base 16 notation
	fmt.Printf("adress of balance in test is %p\n", &wallet.balance)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
