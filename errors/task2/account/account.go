package account

import (
	"errors"
	"math/rand"
)

type Account struct {
	balance int
}

func CreateAccount() *Account {
	return &Account{}
}

func (a *Account) Balance() (int, error) {
	if rand.Intn(10) < 3 {
		return 0, errors.New("Something went wrong")
	}
	return a.balance, nil
}

func (a *Account) Withdraw(amount int) error {
	if amount < 0 {
		return errors.New("Negative amount")
	}
	if amount > a.balance {
		return errors.New("Insufficient funds")
	}

	if rand.Intn(10) < 3 {
		return errors.New("Something went wrong")
	}
	a.balance -= amount

	return nil
}

func (a *Account) Pay(amount int) error {
	if amount < 0 {
		return errors.New("Negative amount")
	}
	if amount > a.balance {
		return errors.New("Insufficient funds")
	}

	if rand.Intn(10) < 3 {
		return errors.New("Something went wrong")
	}

	a.balance -= amount

	return nil
}

func (a *Account) TopUp(amount int) error {
	if amount < 0 {
		return errors.New("Negative amount")
	}
	if rand.Intn(10) < 3 {
		return errors.New("Something went wrong")
	}

	a.balance += amount
	return nil
}
