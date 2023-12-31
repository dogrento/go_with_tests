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

  t.Run("Table driven tests", func(t *testing.T){
    perimeterTests := []struct{
      name string
      shape Shape
      hasPerimeter float64
    }{
      {name: "Rectangle", shape: Rectangle{10, 10}, hasPerimeter: 40},
      {name: "Circle", shape: Circle{10}, hasPerimeter: 62.83185307179586},
      {name: "Triangle", shape: Triangle{Height: 10, Base: 10}, hasPerimeter:30},
    }

    for _, tt := range perimeterTests{
      got := tt.shape.Perimeter()
      if got != tt.hasPerimeter{
        t.Errorf("\n%#v\ngot -> %g\nhasPerimeter -> %g", tt.shape,got, tt.hasPerimeter)
      }
    }
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
      name string
      shape Shape
      hasArea float64
    }{
      {name: "Rectangle",shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
      {name: "Circle",shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
      {name: "Triangle",shape: Triangle{Height: 12, Base: 6}, hasArea: 36.0},
    }

    for _, tt := range areaTests{
      got := tt.shape.Area()
      if got != tt.hasArea{
        t.Errorf("\n%#v\ngot -> %g\nhasArea -> %g", tt.shape,got, tt.hasArea)
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
