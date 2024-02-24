package clockface_test

import(
  "go_with_tests/clockface"
  "testing"
  "time"
)

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
