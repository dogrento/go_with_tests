package main

const PI = 3.141592653589793

type Rectangle struct{
  Width float64
  Height float64
}

// Convention to have to receiver var be the first letter of the typo
// e.g r Rectangle | c Circle
func (r Rectangle) Area() float64{
  return r.Width * r.Height
}

type Circle struct{
  Radius float64
}

func(c Circle) Area() float64{
  return (2 * PI * c.Radius * c.Radius)/2
}

func Perimeter(rec Rectangle) float64{
  result := 2*(rec.Width+rec.Height) 
  return result 
}

func Area(rec Rectangle) float64{
  return rec.Width*rec.Height
}
