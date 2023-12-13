package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T){
  assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin){
    t.Helper()
    got := wallet.Balance()
    msg := fmt.Sprintf("\ngot -> %s\nwant -> %s", got, want)

    if got != want{
      t.Errorf(msg)
    }else{
      fmt.Println(msg)
    }
  }
  assertError := func(t testing.TB, got error, want string){
    t.Helper()
    if got == nil{
      // t.Fatal will stop the test if it is called
      // - without this the test would carry on to tyhe next step and panic because of a nil pointer.
      t.Fatal("seeking error but none found")
    }

    if got.Error() != want{
      t.Errorf("\ngot -> %s\nwant -> %s", got, want)
    }
  }

  t.Run("Get Balance", func(t *testing.T){
    wallet := Wallet{}
    wallet.Deposit(10.0)

    want := Bitcoin(10) 

    assertBalance(t, wallet, want)
  })

  t.Run("withdraw from balance", func(t *testing.T){
    value := Bitcoin(20)
    wallet := Wallet{balance: value}
    wallet.Withdraw(Bitcoin(10))

    want := Bitcoin(10)

    assertBalance(t, wallet, want)
  })

  t.Run("Withdraw insufficient funds", func(t *testing.T){
    startingBalance:= Bitcoin(20)
    withdrawBalance := Bitcoin(100)
    wallet := Wallet{startingBalance} 
    err := wallet.Withdraw(withdrawBalance)

    assertError(t, err, "Cannot withdraw, insufficient funds.")
    assertBalance(t, wallet, startingBalance)
  })
}
