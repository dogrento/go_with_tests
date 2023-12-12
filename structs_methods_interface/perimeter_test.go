package main

import (
	"testing"
)

func TestPerimeter(t *testing.T){
  checkPerimeter := func(t testing.TB, shape Shape, want float64){
    t.Helper()
    got := shape.Perimeter()

    if got != want{
      t.Errorf("\ngot -> %g\nwant -> %g", got, want)
    }
  }
  t.Run("Rectangle perimeter", func(t *testing.T){
    rectangle := Rectangle{10.0, 10.0}
    want := 40.0 

    checkPerimeter(t, rectangle, want)
  })

  t.Run("Circle perimeter", func(t *testing.T){
    circle := Circle{10}
    want := 62.83185307179586

    checkPerimeter(t, circle, want)
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

  t.Run("Table driven tests", func(t *testing.T){
    // Creating anonymous struct
    // - declaring a slice of structs ([]struct)
    // - then fill de slice with cases
    areaTests := []struct{
      shape Shape
      want float64
    }{
      {Rectangle{12, 6}, 72.0},
      {Circle{10}, 314.1592653589793},
    }

    for _, tt := range areaTests{
      got := tt.shape.Area()
      if got != tt.want{
        t.Errorf("\ngot -> %g\nwant -> %g", got, tt.want)
      }
    }
  })
}

func CompareResult(t testing.TB, got float64, want float64){
  t.Helper() 

  if got != want{
    t.Errorf("\ngot -> %.2f\nwant -> %.2f", got, want)
  }
}
