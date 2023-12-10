package main

// Importing packages
import (
	"fmt"
)

const(
  engHelloPrefix = "Hello, "
  spaHelloPrefix = "Hola, "
  frHelloPrefix = "Bonjour, "
  hueHelloPrefix = "Salve, "
                              
  spanish = "spanish"
  french = "french"
  huehue = "huehue"
)

//Creating a function that returns a string
func Hello(name string, lang string) string {
  if name == ""{
    name = "World"
  }

  return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
  switch lang{
    case "french":
      prefix = frHelloPrefix
    case "spanish":
      prefix = spaHelloPrefix
    case "huehue":
      prefix = hueHelloPrefix
    default:
      prefix = engHelloPrefix
    }
  return
}

func main(){
  fmt.Println(Hello("Dgt", ""))
}

