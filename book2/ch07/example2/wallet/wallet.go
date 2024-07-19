package wallet

import (
	"fmt"
)

type Wallet struct {
	balance int
}

// *Wallet - is a pointer to a wallet
func (w *Wallet) Deposit(amount int) {
	// %p prints memory address in base 16 notation
	fmt.Printf("adress of balance in Deposit is %p\n", &w.balance)
	// or: (*w).balance += amount
	w.balance += amount
}

// *Wallet - is a pointer to a wallet
func (w *Wallet) Balance() int {
	// or: return (*w).balance
	return w.balance
}
