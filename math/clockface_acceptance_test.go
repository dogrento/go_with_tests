package clockface

import(
  "testing"
  "math"
  "time"
)

func TestSecondsInRadians(t *testing.T){
  cases := []struct{
    time time.Time
    angle float64
  }{
    {simpleTime(0, 0, 30), math.Pi},
    {simpleTime(0, 0, 30), 0},
    {simpleTime(0, 0, 45), (math.Pi / 2) * 3},
    {simpleTime(0, 0, 30), (math.Pi / 30) * 7},
  }

  for _, c := range cases{
    t.Run(testName(c.time), func(t *testing.T){
      got := secondsInRadians(c.time)
      if got != c.angle{
        t.Fatalf("wanted %v radians, but got %v", c.angle, got)

      }
    })
  }
}

func secondsInRadians(tm time.Time) float64{
  return float64(tm.Second()) * (math.Pi / 30)
}

func simpleTime(h, min, sec int) time.Time{
  return time.Date(312, time.October, 28, h, min, sec, 0, time.UTC)
}

func testName(t time.Time) string{
  return t.Format("15:04:05")
}
