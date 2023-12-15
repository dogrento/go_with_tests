package main

import(
  "testing"
  "bytes"
)

func TestCountdown(t *testing.T){
  buffer := &bytes.Buffer{}

  Countdown(buffer)

  got := buffer.String()
  want := "3"

  if got != want{
    t.Errorf("\ngot->%s\nwant->%s", got, want)
  }
}
