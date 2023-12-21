package main

import(
  "testing"
)

func TestRomanNumerals(t *testing.T){
  t.Run("1 gets converted to I", func(t *testing.T){
    got := ConvertToRoman(1)
    want := "I"

    // if got != want{
    //   t.Errorf("got: %q, want: %q", got, want)
    // }
    assertRomanNumerals(t, got, want)
  })

  t.Run("2 gets converted to II", func(t *testing.T){
    got := ConvertToRoman(2)
    want := "II"

    assertRomanNumerals(t, got, want)
  })
}

func assertRomanNumerals(t testing.TB, got, want string){
  t.Helper()
  if got != want{
    t.Errorf("got: %q, want: %q", got, want)
  }
}
