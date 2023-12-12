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

func Perimeter(rec Rectangle) float64{
  result := 2*(rec.Width+rec.Height) 
  return result 
}

// func Area(rec Rectangle) float64{
//   return rec.Width*rec.Height
// }
