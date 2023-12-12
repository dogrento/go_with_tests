package main

import (
	"testing"
)

func TestPerimeter(t *testing.T){
  t.Run("Rectangle perimeter result", func(t *testing.T){
    rectangle := Rectangle{10.0, 10.0}

    got := Perimeter(rectangle)
    want := 40.0 

    CompareResult(t, got, want)
  })
}

func TestArea(t *testing.T){

  checkArea := func(t testing.TB, shape Shape, want float64){
    t.Helper()
    got := shape.Area()
    if got != want{
      t.Errorf("\ngot -> %g\nwant -> %g", got, want)
    }
  }

  t.Run("Rectangle area", func(t *testing.T){
    rectangle := Rectangle{12.0, 6.0}
    want := 72.0

    checkArea(t, rectangle, want)
  })

  t.Run("Circle area", func(t *testing.T){
    var rad float64 
    rad = 10
    circle := Circle{rad}

    want := 314.1592653589793

    checkArea(t, circle, want)
  })
}

func CompareResult(t testing.TB, got float64, want float64){
  t.Helper() 

  if got != want{
    t.Errorf("\ngot -> %.2f\nwant -> %.2f", got, want)
  }
}
