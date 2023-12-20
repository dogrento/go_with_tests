package main

import (
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, store.Fetch())
  }
}
