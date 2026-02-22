package main

type AccountRepository interface {
	GetAccount(id string) (*Account, error)
	UpdateAccount(acc *Account) error
}
