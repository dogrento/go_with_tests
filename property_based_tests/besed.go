package main

import(
  "strings"
)

type RomanNumeral struct{
  Value  int
  Symbol string
}

var allRomanNumerals = []RomanNumeral{
  {50, "L"},
  {40, "XL"},
  {10, "X"},
  {9, "IX"},
  {5, "V"},
  {4, "IV"},
  {1, "I"},
}

func ConvertToRoman(value int) string{

  // Builder is used to efficiently build a string using Write methods
  // - It minimizes memory copying
  var result strings.Builder

  for _, numeral := range allRomanNumerals{
    for value >= numeral.Value{
      result.WriteString(numeral.Symbol)
      value -= numeral.Value
    }
  }

  return result.String()
}
