package main

import(
  "testing"
)

func TestRacer(t *testing.T){
  t.Run("racer test", func(t *testing.T){
    slowUrl := "http://www.facebook.com"
    fastUrl := "http://www.quii.dev"

    want := fastUrl
    got := Racer(slowUrl, fastUrl)

    if got != want{
      t.Errorf("\ngot -> %s\n,want -> %s", got, want)
    }
  })
}
