package greet

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(w io.Writer, name string){
  fmt.Fprintf(w, "Hello, %s", name)
}

// When you write an HTTP handler, you are given an http.ResponseWriter and the http.Request
// that was used to make the resquest.
// When you implement your server, you WRITE your repsonse using the writer
func MyGreeterHandler(w http.ResponseWriter, r *http.Request){
  Greet(w, "world")
}
