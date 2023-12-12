package main

import "math"

type Shape interface{
  Area() float64
  Perimeter() float64
}

type Rectangle struct{
  Width float64
  Height float64
}
// Convention to have to receiver var be the first letter of the typo
// e.g r Rectangle | c Circle
func (r Rectangle) Area() float64{
  return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64{
  return 2 * (r.Height + r.Width)
}

type Circle struct{
  Radius float64
}
func(c Circle) Area() float64{
  return (2 * math.Pi * c.Radius * c.Radius)/2
}
func(c Circle) Perimeter() float64{
  return 2 * math.Pi * c.Radius
}

type Triangle struct{
  Height float64
  Base float64
}
func(t Triangle) Area() float64{
  return (t.Base * t.Height) / 2 //or (t.Base * t.Height) * 0.5
}
func(t Triangle) Perimeter() float64{
  return t.Base * 3  
}
