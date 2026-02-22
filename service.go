package main

import "errors"

//defines business logic operations
type TransferService interface {
	SendMoney(fromID string, toID string, amount float64) error
}

// Dependency Injection - holds our repo interface!!
type BankService struct {
	repo AccountRepository
}

//The Injector
func NewBankService(r AccountRepository) *BankService {
	return &BankService{
		repo: r,
	}
}

func (s *BankService) SendMoney(fromID string, toID string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	sender, err := s.repo.GetAccount(fromID)
	if err != nil {
		return errors.New("sender account does not exist")
	}

	receiver, err := s.repo.GetAccount(toID)
	if err != nil {
		return errors.New("receiver account doesnt exist")
	}

	if sender.Balance < amount {
		return errors.New("insufficient funds")
	}

	sender.Balance -= amount
	receiver.Balance += amount

	//save them back to DB - Sender
	err = s.repo.UpdateAccount(sender)
	if err != nil {
		return err
	}
	//save them back to DB - Receiver
	err = s.repo.UpdateAccount(receiver)
	if err != nil {
		return err
	}

	//If success (all cases have passed)
	return nil

}
