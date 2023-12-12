package main

import (
  "testing"
)

func TestWallet(t *testing.T){
  assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin){
    t.Helper()
    got := wallet.Balance()

    if got != want{
      t.Errorf("\ngot -> %s\nwant -> %s", got, want)
    }
  }
  assertError := func(t testing.TB, err error){
    t.Helper()
    if err == nil{
      t.Error("seeking error but none found")
    }
  }

  t.Run("Get Balance", func(t *testing.T){
    wallet := Wallet{}
    wallet.Deposit(10.0)

    want := Bitcoin(10) 

    assertBalance(t, wallet, want)
  })

  t.Run("withdraw from balance", func(t *testing.T){
    wallet := Wallet{balance: Bitcoin(20)}
    wallet.Withdraw(Bitcoin(10))

    want := Bitcoin(10)

    assertBalance(t, wallet, want)
  })

  t.Run("Withdraw insufficient funds", func(t *testing.T){
    startingBalance:= Bitcoin(20)
    withdrawBalance := Bitcoin(100)
    wallet := Wallet{startingBalance} 
    err := wallet.Withdraw(withdrawBalance)

    assertBalance(t, wallet, startingBalance)
    assertError(t, err)

    // if err == nil{
    //   t.Error("seeking error but none found")
    // }
  })
}
