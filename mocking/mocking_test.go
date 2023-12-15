package main

import(
  "testing"
  "bytes"
)

func TestCountdown(t *testing.T){
  buffer := &bytes.Buffer{}

  Countdown(buffer)

  got := buffer.String()
  // backtick is another way of creating a string but lets you include things like newlines
  want := `3
2
1
Go!`

  if got != want{
    t.Errorf("\ngot->%s\nwant->%s", got, want)
  }
}
