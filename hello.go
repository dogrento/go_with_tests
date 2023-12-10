package main

// Importing packages
import "fmt"

//Creating a function that returns a string
func Hello(name string) string {
  return "Hello " + name
}

func main(){
  fmt.Println(Hello("Dgt"))
}
