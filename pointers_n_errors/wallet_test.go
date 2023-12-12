package main

import (
  "testing"
  "fmt"
)

func TestWallet(t *testing.T){
  t.Run("Get Balance", func(t *testing.T){
    wallet := Wallet{}
    wallet.Deposit(10.0)

    got := wallet.Balance()
    fmt.Printf("addr of balance in Test -> %p\n", &wallet.balance)
    want := 10.0

    if got != want{
      t.Errorf("\ngot -> %g\nwant -> %g", got, want)
    }
  })
}
