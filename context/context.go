package main

import (
	"fmt"
	"net/http"
)

type Store interface{
  Fetch() string
  Cancel()
}

func Server(store Store) http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request){
    ctx := r.Context()

    data := make(chan string, 1)

    go func(){
      // Fetch() will write the result into a new chan 
      data <- store.Fetch()
    }()

    // Using select to effectivbely race to the two asynchronous processes and then 
    // either write a response or cancel.
    select{
    case d := <-data:
      fmt.Fprint(w, d)
    //Done() returns a chan which gets sent a signal when the context is done or cancelled.
    case <-ctx.Done():
      store.Cancel()
    }
  }
}
