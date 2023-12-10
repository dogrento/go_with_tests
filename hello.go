package main

// Importing packages
import "fmt"

const engHelloPrefix = "Hello, "
const spaHelloPrefix = "Hola, "
const frHelloPrefix = "Bonjour, "
const hueHelloPrefix = "Salve, "

const spanish = "spanish"
const french = "french"
const huehue = "huehue"

//Creating a function that returns a string
func Hello(name string, lang string) string {
  if name == ""{
    name = "World"
  }

  prefix := engHelloPrefix

  switch lang{
  case "french":
    prefix = frHelloPrefix
  case "spanish":
    prefix = spaHelloPrefix
  case "huehue":
    prefix = hueHelloPrefix
  }

  return prefix + name
}

func main(){
  fmt.Println(Hello("Dgt", ""))
}
