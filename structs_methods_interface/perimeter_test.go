package main

import (
	"testing"
)

func TestPerimeter(t *testing.T){
  t.Run("Perimeter result", func(t *testing.T){
    width := 10.0
    height := 10.0

    got := Perimeter(width, height)
    want := 40.0 

    if got != want{
      t.Errorf("\ngot -> %.2f\nwant -> %.2f", got, want)
    }
  })
}
