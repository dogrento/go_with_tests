package main

import(
  "strings"
)

func ConvertToRoman(value int) string{

  // Builder is used to efficiently build a string using Write methods
  // - It minimizes memory copying
  var result strings.Builder

  for i := value; i > 0; i--{
    if i == 4{
      result.WriteString("IV")
      break
    }
    if i == 5{
      result.WriteString("V")
      break
    }
    result.WriteString("I")
  }

  return result.String()
}
