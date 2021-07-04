package wallet

import (
	"errors"
	"fmt"
)

var InsufficientFundsError = errors.New("取的钱比账户余额还要多！！")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	} else {
		w.balance -= amount
	}

	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
