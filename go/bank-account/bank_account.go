package account

import "sync"

// Account represents a bank account.
type Account struct {
	m       sync.Mutex
	balance int64
	open    bool
}

// Open creates a new bank account with an inital balance equal to
// initialDeposit.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{sync.Mutex{}, initialDeposit, true}
}

// Close closes a bank account. Ok is false if the account is already closed.
func (a *Account) Close() (payout int64, ok bool) {
	a.m.Lock()
	defer a.m.Unlock()

	if !a.open {
		return
	}

	ok = true
	payout = a.balance
	a.open = false
	return
}

// Balance returns the ammount of money deposited in the account. Ok is false
// if the account is closed.
func (a *Account) Balance() (balance int64, ok bool) {
	a.m.Lock()
	defer a.m.Unlock()

	if !a.open {
		return
	}

	ok = true
	balance = a.balance
	return
}

// Deposit adds amount to the balance of the account. If ammount is negative
// a withdrawal is made. Ok is false if either the account is closed or a
// withdrawal exceeds the balance.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.m.Lock()
	defer a.m.Unlock()

	if !a.open || a.balance+amount < 0 {
		return
	}

	ok = true
	a.balance += amount
	newBalance = a.balance
	return
}
