package main

import (
  "go_with_tests/dependency_injection"
  "log"
  "net/http"
)

func main(){
  // greet.Greet(os.Stdout, "Elodie")
  log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(greet.MyGreeterHandler)))
}
