package main

import(
  "testing"
)

func TestRomanNumerals(t *testing.T){
  cases := []struct{
    Description string
    Arabic      int
    Want        string
  }{
    {"1 gets converted to I", 1, "I"},
    {"2 gets converted to II", 2, "II"},
    {"3 gets converted to III", 3, "III"},
  }

  for _, test := range cases{
    t.Run(test.Description, func(t *testing.T){
      got := ConvertToRoman(test.Arabic)
      assertRomanNumerals(t, got, test.Want)
    })
  }
}

func assertRomanNumerals(t testing.TB, got, want string){
  t.Helper()
  if got != want{
    t.Errorf("got: %q, want: %q", got, want)
  }
}
