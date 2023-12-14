package main

import(
  "testing"
  "bytes"  
)

func TestGreet(t *testing.T){
  t.Run("testing greet", func(t *testing.T){
    buffer := bytes.Buffer{}
    Greet(&buffer, "Dgt")
    
    got := buffer.String()
    want := "Hello, Dgt"

    if got != want{
      t.Errorf("\ngot -> %s\nwant -> %s", got, want)
    }

  })
}
