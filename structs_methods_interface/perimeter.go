package main

type Rectangle struct{
  Width float64
  Height float64
}

func Perimeter(rec Rectangle) float64{
  result := 2*(rec.Width+rec.Height) 
  return result 
}

func Area(rec Rectangle) float64{
  return rec.Width*rec.Height
}
