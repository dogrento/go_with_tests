package main

import (
	// "bytes"
	"fmt"
	"io"
	"os"
  "time"
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
    time.Sleep(1 * time.Second)
  }
  fmt.Fprint(out, finalWord)
}

func main(){
  Countdown(os.Stdout)
}
