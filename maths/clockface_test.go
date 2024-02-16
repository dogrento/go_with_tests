package clockface_test

import (
	"go_with_tests/maths"
	"strings"
	"testing"
	"time"
)

func TestSecondHandPoint(t *testing.T){
  cases := []struct{
    time time.Time
    point clockface.Point
  }{
    {simpleTime(0, 0, 30), clockface.Point{0, -1}},
  }

  for _, c := range cases{
    t.Run(testName(c.time), func(t *testing.T){
      got := secondHandPoint(c.time)
      if got != c.point{
        t.Fatalf("wanted %v Point, but got %v", c.point, got)
      }
    })
  }
}

func TestSecondHandAtMidnight(t *testing.T){
  tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

  want := clockface.Point{X: 150, Y: 150 - 90}
  got := clockface.SecondHand(tm)

  if got != want{
    t.Errorf("got %v, wanted %v", got, want)
  }
}

func secondHandPoint(tm time.Time) clockface.Point{
  return clockface.Point{}
}

func TestSVGWriterAtMidnight(t *testing.T){
  tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

  var b strings.Builder
  clockface.SVGWriter(&b, tm)
  got := b.String()

  want := `<line x1="150" y1="150" x2="150" y2="60"`

  if !strings.Contains(got, want){
    t.Errorf("Expected to find the second hand %v, in the SVG output %v", want, got)
  }
}

// func TestSecondHandAt30Seconds(t *testing.T){
//   tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

//   want := clockface.Point{X: 150, Y: 150 + 90}
//   got := clockface.SecondHand(tm)

//   if got != want{
//     t.Errorf("got %v, wanted %v", got, want)
//   }
// }

func simpleTime(h, min, sec int) time.Time{
  return time.Date(312, time.October, 28, h, min, sec, 0, time.UTC)
}

func testName(t time.Time) string{
  return t.Format("15:04:05")
}
