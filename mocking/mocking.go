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

// Creating a real sleeper which uses the interface
type DefaultSleeper struct{}
func (d *DefaultSleeper) Sleep(){
  time.Sleep(1 * time.Second)
}

const(
  finalWord = "Go!"
  countdownStart = 3
)

// it takes an io.Writer
// and sends a string
// func Countdown(out *bytes.Buffer){
func Countdown(out io.Writer, sleeper Sleeper){
  for i := countdownStart; i > 0; i--{
    fmt.Fprintln(out, i)
    sleeper.Sleep()
  }

  // for i := countdownStart; i >0; i--{
  //   fmt.Fprintln(out, i)
  // }
  fmt.Fprint(out, finalWord)
}

func main(){
  sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
  Countdown(os.Stdout, sleeper)
}
