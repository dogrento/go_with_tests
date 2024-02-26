package clockface

import(
  "time"
  "math"
)

// Point represents a 2d cartesian coordinate
type Point struct{
  X float64
  Y float64
}

func secondHandPoint(t time.Time) Point{
  return angleToPoint(secondsInRadians(t))
}

func secondsInRadians(t time.Time) float64{
  return (math.Pi / (30 / (float64(t.Second()))))
}

func minuteHandPoint(t time.Time) Point{
  return angleToPoint(minutesInRadians(t))
}

func minutesInRadians(t time.Time) float64{
  return (secondsInRadians(t) / 60 + (math.Pi / (30 / float64(t.Minute())))) 
}

func angleToPoint(angle float64) Point{
  x := math.Sin(angle)
  y := math.Cos(angle)

  return Point{x, y}
}

