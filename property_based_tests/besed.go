package main

import(
  "strings"
)

func ConvertToRoman(value int) string{
  var result strings.Builder

  for i := 0; i < value; i++{
    result.WriteString("I")
  }

  return result.String()
}
