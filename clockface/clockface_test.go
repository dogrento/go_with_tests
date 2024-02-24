package clockface_test

import(
  "go_with_tests/clockface"
  "testing"
  "time"
  "math"
)

func TestSecondsInRadians(t *testing.T){
  thirtySec := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
  want := math.Pi
  got := secondsInRadians(thirtySec)

  if got != want{
    t.Errorf("got %v, wanted %v", got, want)
  }
}

func secondsInRadians(t time.Time) float64{
  return math.Pi
}

func TestSecondHandAtMidnight(t *testing.T){
  tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

  want := clockface.Point{
    X: 150,
    Y: 150 - 90,
  }

  got := clockface.SecondHand(tm)

  if got != want{
    t.Errorf("got %v, wanted %v", got, want)
  }
}

func TestSecondHandAt30Sec(t *testing.T){
  tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

  want := clockface.Point{
    X: 150,
    Y: 150 + 90,
  }

  got := clockface.SecondHand(tm)

  if got != want{
    t.Errorf("got %v, wanted %v", got, want)
  }
}
