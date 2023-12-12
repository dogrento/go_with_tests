package main

import (
  "testing"
)

func TestWallet(t *testing.T){
  t.Run("Get Balance", func(t *testing.T){
    wallet := Wallet{}
    wallet.Deposit(10.0)

    got := wallet.Balance()
    want := Bitcoin(10) 

    if got != want{
      t.Errorf("\ngot -> %g\nwant -> %g", got, want)
    }
  })
}
