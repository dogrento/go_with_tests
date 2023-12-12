package main

import (
	"fmt"
  "errors"
)

// Creating new types from existing ones
// - to make Bitcoin type, use syntax: Bitcoin(value)
type Bitcoin float64

type Stringer interface{
  String() string
}
func (b Bitcoin) String() string{
  return fmt.Sprintf("%g BTC", b)
}

type Wallet struct{
  // if a symbol starts with a lowercase then it's private outside the package
  balance Bitcoin
}
// The reason to use pointer is bc the var is a copy when it is passed as param
// 'struct pointer' are automatically dereferenced
func(w *Wallet) Deposit(value Bitcoin){
  w.balance += value
}
func(w *Wallet) Balance() Bitcoin{
  return w.balance
}
func(w *Wallet) Withdraw(value Bitcoin) error{
  if value > w.balance{
    return errors.New("Oh, snap!")
  }
  w.balance -= value
  return nil
}
