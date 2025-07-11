package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

// When %s is used to print any Bitcoin, is gonna call the String() function
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// The pointer gets the original value of the wallet called by the test, not a copy
func (w *Wallet) Deposit(amount Bitcoin) {
	// No need to put like *(w).balance, go deals this way
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrorInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrorInsufficientFunds
	}
	w.balance -= amount
	return nil
}
