package pointerserrors

import "errors"

type Coin int

type Wallet struct {
	balance Coin
}

var ErrInsufficientFunds = errors.New("Cannot Withdrawal Insufficient Funds")

func (w *Wallet) Deposit(amount Coin) {
	w.balance += amount
}

func (w *Wallet) Balance() Coin {
	return w.balance
}

func (w *Wallet) Withdrawal(amount Coin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
