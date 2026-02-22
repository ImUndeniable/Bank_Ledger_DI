package main

import (
	"errors"
	"sync"
)

type MemoryRepo struct {
	accounts map[string]*Account
	mu       sync.RWMutex
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		accounts: map[string]*Account{
			"alice": {ID: "alice", Balance: 125.50},
			"bob":   {ID: "bob", Balance: 50.50},
		},
	}
}

func (m *MemoryRepo) GetAccount(id string) (*Account, error) {
	m.mu.RLock()
	defer m.mu.Unlock()

	account, exists := m.accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}

	return account, nil
}

func (m *MemoryRepo) UpdateAccount(acc *Account) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, exists := m.accounts[acc.ID]

	if !exists {
		return errors.New("account not found")
	} else {
		m.accounts[acc.ID] = acc
	}
	return nil
}
