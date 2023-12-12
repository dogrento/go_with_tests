package main

import "fmt"

type Wallet struct{
  // if a symbol starts with a lowercase then it's private outside the package
  balance float64
}
func(w Wallet) Deposit(value float64){
  fmt.Printf("addr of balance in Deposit() -> %p\n", &w.balance)
  w.balance += value
}
func(w Wallet) Balance() float64{
  return w.balance
}
