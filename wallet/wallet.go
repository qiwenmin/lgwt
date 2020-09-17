package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin is the type of bitcoin
type Bitcoin int

// Stringer can generate a string via String()
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet defines a bitcoin wallet.
type Wallet struct {
	balance Bitcoin
}

// Deposit deposits the bitcoins into the wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds - balance too low to withdraw.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw withdraws the bitcoins from the wallet.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance returns the balance of the wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
