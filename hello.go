package main

// Importing packages
import "fmt"

const helloPrefix = "Hello, "

//Creating a function that returns a string
func Hello(name string) string {
  if name == ""{
    name = "World"
  }
  return helloPrefix + name
}

func main(){
  fmt.Println(Hello("Dgt"))
}
