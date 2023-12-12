package main

// Creating new types from existing ones
type Bitcoin float64

type Wallet struct{
  // if a symbol starts with a lowercase then it's private outside the package
  balance Bitcoin
}
// The reason to use pointer is bc the var is a copy when it is passed as param
// 'struct pointer' are automatically dereferenced
func(w *Wallet) Deposit(value float64){
  w.balance += value
}
func(w *Wallet) Balance() float64{
  return w.balance
}
