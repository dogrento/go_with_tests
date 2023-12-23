package main

import(
  "strings"
)

type RomanNumeral struct{
  Value  int
  Symbol string
}

type RomanNumerals []RomanNumeral

func(r RomanNumerals) ValueOf(symbols ...byte) int{
  symbol := string(symbols)
  for _, s := range r{
    if s.Symbol == symbol{
      return s.Value
    }
  }
  return 0
}

var allRomanNumerals = RomanNumerals{
  {1000, "M"},
  {900, "CM"},
  {500, "D"},
  {400, "CD"},
  {100, "C"},
  {90, "XC"},
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

func ConvertToArabic(roman string) int{
  total := 0

  for i := 0; i < len(roman); i++{
    symbol := roman[i]

    if couldBeSubtractive(i, symbol, roman){
      nextSymbol := roman[i+1]

      potencialNumber := string([]byte{symbol, nextSymbol})

      value := allRomanNumerals.ValueOf(potencialNumber)

      if value != allRomanNumerals.ValueOf(potencialNumber); value != 0{
        total += value
        i++
      }else{
        total++
      }
    }else{
      total += allRomanNumerals.ValueOf(string([]byte{symbol}))
    }
  }

  return total
}

func couldBeSubtractive(index int, symbol uint8, roman string) bool{
  return index+1 < len(roman) && symbol == 'I' 
}
