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

func secondHandPoint(t time.Time) clockface.Point{
  return angleToPoint(secondsInRadians(t))
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

func TestMinutesInRadians(t *testing.T){
  cases := []struct{
    time time.Time
    angle float64
  }{
    {simpleTime(0, 30, 0), math.Pi},
    {simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
  }

  for _, c := range cases{
    t.Run(testName(c.time), func(t *testing.T){
      got := minutesInRadians(c.time)
      if got != c.angle{
        t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
      }
    })
  }
}

func TestMinuteHandPoint(t *testing.T){
  cases := []struct {
    time time.Time
    point clockface.Point
  }{
    {simpleTime(0, 30, 0), clockface.Point{0, -1}},
  }

  for _, c := range cases{
    t.Run(testName(c.time), func(t *testing.T){
      got := minuteHandPoint(c.time)
      if !roughlyEqualPoint(got, c.point){
        t.Fatalf("got %v radians, wanted %v", got, c.point)
      }
    })
  }
}

func minuteHandPoint(t time.Time) clockface.Point{
  return angleToPoint(minutesInRadians(t))
}

func minutesInRadians(t time.Time) float64{
  return (secondsInRadians(t) / 60 + (math.Pi / (30 / float64(t.Minute())))) 
}

func TestHourInRadians(t *testing.T){
  cases := []struct{
    time time.Time
    angle float64
  }{
    {simpleTime(6, 0, 0), math.Pi},
    {simpleTime(0, 0, 0), 0},
    {simpleTime(21, 0, 0), math.Pi * 1.5},
  }

  for _, c := range cases{
    t.Run(testName(c.time), func(t *testing.T){
      got := hoursInRadians(c.time)
      if got != c.angle{
        t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
      }
    })
  }
}

func hoursInRadians(t time.Time) float64{
  return (math.Pi / (6 / float64(t.Hour())))
}

// func TestHourHandPoint(t *testing.T){
//   cases := []struct {
//     time time.Time
//     point clockface.Point
//   }{
//     {simpleTime(30, 0, 0), clockface.Point{0, -1}},
//   }

//   for _, c := range cases{
//     t.Run(testName(c.time), func(t *testing.T){
//       got := minuteHandPoint(c.time)
//       if !roughlyEqualPoint(got, c.point){
//         t.Fatalf("got %v radians, wanted %v", got, c.point)
//       }
//     })
//   }
// }

func angleToPoint(angle float64) clockface.Point{
  x := math.Sin(angle)
  y := math.Cos(angle)

  return clockface.Point{x, y}
}

func roughlyEqualFloat64(a, b float64) bool{
  const equalityThreshold = 1e-7
  return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b clockface.Point) bool{
  return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func simpleTime(hours, minutes, seconds int) time.Time{
  return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string{
  return t.Format("15:04:05")
}

