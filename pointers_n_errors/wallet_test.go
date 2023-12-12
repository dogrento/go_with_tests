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
      t.Errorf("\ngot -> %s\nwant -> %s", got, want)
    }
  })
  t.Run("withdraw from balance", func(t *testing.T){
    wallet := Wallet{balance: Bitcoin(20)}

    wallet.Withdraw(Bitcoin(10))

    got := wallet.Balance()
    want := Bitcoin(10)

    if got != want{
      t.Errorf("\ngot -> %s\nwant -> %s", got, want)
    }
    
  })
}
