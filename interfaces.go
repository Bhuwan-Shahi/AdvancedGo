package main

import (
	"errors"
	"fmt"
)

type bankAcc interface {
	GetBalance() int
	Deposit(amt int)
	WithDraw(amt int) error
}

type siddartha struct {
	balance int
}

// constructor
func newSiddartha() *siddartha {
	return &siddartha{
		balance: 0,
	}
}

func (s *siddartha) GetBalance() int {
	return s.balance
}

func (s *siddartha) Deposit(amt int) {
	s.balance += amt
	fmt.Println(s.balance, ",This is you balance")

}

func (s *siddartha) WithDraw(amt int) error {
	newBalance := s.balance - amt
	if newBalance < 0 {
		return errors.New("Insufficient funds")
	}
	s.balance = newBalance
	return nil
}

func main() {
	sid := siddartha{
		balance: 0,
	}
	sid.Deposit(345)
	fmt.Println(sid.GetBalance())
	fmt.Println(sid.WithDraw(567))
}
