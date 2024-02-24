package clockface_test

import(
  "go_with_tests/clockface"
  "testing"
  "time"
  "math"
)

func TestSecondsInRadians(t *testing.T){
  cases := []struct{
    time time.Time
    angle float64
  }{
    {simpleTime(0, 0, 30),           math.Pi},
    {simpleTime(0, 0, 0),                  0},
    {simpleTime(0, 0, 45), (math.Pi / 2) * 3},
    {simpleTime(0, 0, 7), (math.Pi / 30) * 7},
  }

  for _, c := range cases{
    t.Run(testName(c.time), func(t *testing.T){
      got := secondsInRadians(c.time)
      if got != c.angle{
        t.Fatalf("got %v radians, wanted %v", got, c.angle)
      }
    })
  }
}

func secondsInRadians(t time.Time) float64{
  return (math.Pi / (30 / (float64(t.Second()))))
}

func TestSecondHandPoint(t *testing.T){
  cases := []struct {
    time time.Time
    point clockface.Point
  }{
    {simpleTime(0, 0, 30), clockface.Point{0, -1}},
    {simpleTime(0, 0, 45), clockface.Point{-1, 0}},
  }

  for _, c := range cases{
    t.Run(testName(c.time), func(t *testing.T){
      got := secondHandPoint(c.time)
      if !roughlyEqualPoint(got, c.point){
        t.Fatalf("got %v radians, wanted %v", got, c.point)
      }
    })
  }
}

func roughlyEqualFloat64(a, b float64) bool{
  const equalityThreshold = 1e-7
  return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b clockface.Point) bool{
  return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func secondHandPoint(t time.Time) clockface.Point{
  angle := secondsInRadians(t)
  x:= math.Sin(angle)
  y := math.Cos(angle)
  return clockface.Point{x, y}
}


func simpleTime(hours, minutes, seconds int) time.Time{
  return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string{
  return t.Format("15:04:05")
}

