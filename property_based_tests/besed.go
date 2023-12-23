package main

import(
  "strings"
)

type RomanNumeral struct{
  Value  uint16
  Symbol string
}

type RomanNumerals []RomanNumeral

func(r RomanNumerals) ValueOf(symbols ...byte) uint16{
  symbol := string(symbols)
  for _, s := range r{
    if s.Symbol == symbol{
      return s.Value
    }
  }
  return 0
}

func(r RomanNumerals) Exists(symbols ...byte) bool{
  symbol := string(symbols)
  for _, s := range r{
    if s.Symbol == symbol{
      return true
    }
  }
  return false
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

type windowedRoman string
func(w windowedRoman) Symbols() (symbols [][]byte){
  for i := 0; i < len(w); i++{
    symbol := w[1]
    notAtEnd := i+1 < len(w)

    if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]){
      symbols = append(symbols, []byte{symbol, w[i+1]})
      i++
    }else{
      symbols = append(symbols, []byte{symbol})
    }
  }
  return
}

func isSubtractive(symbol uint8) bool{
  return symbol == 'I' || symbol == 'X' || symbol == 'C'
}

func ConvertToRoman(value uint16) string{

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

func ConvertToArabic(roman string) (total uint16){
  for _, symbols := range windowedRoman(roman).Symbols(){
    total += allRomanNumerals.ValueOf(symbols...)
  }
  return 
}

// func couldBeSubtractive(index uint16, symbol uint16, roman string) bool{
//   isSubtractiveSymbol := symbol == 'I' || symbol == 'X' || symbol == 'C'
//   return index+1 < len(roman) && isSubtractiveSymbol
// }
