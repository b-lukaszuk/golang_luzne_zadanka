package wallet

import (
	"fmt"
)

type Wallet struct {
	balance int
}

// operates on a copy of a wallet
func (w Wallet) Deposit(amount int) {
	// %p prints memory address in base 16 notation
	fmt.Printf("adress of balance in Deposit is %p\n", &w.balance)
	w.balance += amount
}

// operates on a copy of a wallet
func (w Wallet) Balance() int {
	return w.balance
}
