package main

import (
	// "bytes"
	"fmt"
	"io"
	"os"
  "time"
)

type Sleeper interface{
  Sleep()
}

//MOCK
// Spies are a kind of mock which can record how a dependency is used
// they can record the arguments sent in, how many times it has been called, etc.
// In this case, this gonna track of how many time Sleep() is called
type SpySleeper struct{
  Calls int
}

func (s *SpySleeper) Sleep(){
  s.Calls++
}

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
