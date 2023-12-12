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

    // if got != want{
    //   t.Errorf("\ngot -> %.2f\nwant -> %.2f", got, want)
    // }
    CompareResult(t, got, want)
  })
}

func TestArea(t *testing.T){
  t.Run("Area Result", func(t *testing.T){
    got := Area(12.0, 6.0)
    want := 72.0

    CompareResult(t, got, want)
  })
}

func CompareResult(t testing.TB, got float64, want float64){
  t.Helper() 

  if got != want{
    t.Errorf("\ngot -> %.2f\nwant -> %.2f", got, want)
  }
}
