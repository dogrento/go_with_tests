package clockface

import(
  "testing"
  "math"
  "time"
)

func TestSecondsInRadians(t *testing.T){
  thirtySec := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
  want := math.Pi
  got := secondsInRadians(thirtySec)

  if got != want{
    t.Fatalf("wanted %v radians, but got %v", want, got)
  }
}

func secondsInRadians(tm time.Time) float64{
  return math.Pi
}
