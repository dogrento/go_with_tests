package main

import(
  "strings"
)

func ConvertToRoman(value int) string{

  // Builder is used to efficiently build a string using Write methods
  // - It minimizes memory copying
  var result strings.Builder

  for value > 0{
    switch{
    case value > 8:
      result.WriteString("IX")
      value -= 9
    case value > 4:
      result.WriteString("V")
      value -= 5
    case value > 3:
      result.WriteString("IV")
      value -= 4
    default:
      result.WriteString("I")
      value--
    }
  }

  return result.String()
}
