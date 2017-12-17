package account

import (
	"sync"
)

type Account struct {
	balance int64
	closed  bool
	mutex   *sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{initialDeposit, false, &sync.Mutex{}}
}

func (acc *Account) Close() (payout int64, ok bool) {
	acc.mutex.Lock()
	if acc.closed {
		return 0, false
	}
	acc.closed = true
	acc.mutex.Unlock()

	payout = acc.balance
	ok = true
	return
}

func (acc Account) Balance() (balance int64, ok bool) {
	if acc.closed {
		return 0, false
	}
	return acc.balance, true
}

func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	if acc.closed {
		return 0, false
	}
	acc.mutex.Lock()
	defer acc.mutex.Unlock()
	newBalance = acc.balance + amount
	if newBalance < 0 {
		newBalance = acc.balance
		ok = false
	} else {
		acc.balance = newBalance
		ok = true
	}
	return
}
