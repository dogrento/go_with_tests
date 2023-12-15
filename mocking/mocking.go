package main

import (
	// "bytes"
	"fmt"
	"io"
	"os"
)

const(
  finalWord = "Go!"
  countdownStart = 3
)


// it takes an io.Writer
// and sends a string
// func Countdown(out *bytes.Buffer){
func Countdown(out io.Writer){
  for i := countdownStart; i > 0; i--{
    fmt.Fprintln(out, i)
  }
  fmt.Fprint(out, finalWord)
}

func main(){
  Countdown(os.Stdout)
}
