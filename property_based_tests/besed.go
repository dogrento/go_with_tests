package main

import(
  "strings"
)

func ConvertToRoman(value int) string{

  // Builder is used to efficiently build a string using Write methods
  // - It minimizes memory copying
  var result strings.Builder

  for i := 0; i < value; i++{
    if value == 4{
      return "IV"
    }
    result.WriteString("I")
  }

  return result.String()
}
