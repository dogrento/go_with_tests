package clockface_test

import (
	"bytes"
	"encoding/xml"
	"go_with_tests/clockface"
	"testing"
	"time"
)

// func TestSecondHandAtMidnight(t *testing.T){
//   tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

//   want := clockface.Point{
//     X: 150,
//     Y: 150 - 90,
//   }

//   got := clockface.secondHandPoint(tm)

//   if got != want{
//     t.Errorf("got %v, wanted %v", got, want)
//   }
// }

// func TestSecondHandAt30Sec(t *testing.T){
//   tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

//   want := clockface.Point{
//     X: 150,
//     Y: 150 + 90,
//   }

//   got := clockface.secondHand(tm)

//   if got != want{
//     t.Errorf("got %v, wanted %v", got, want)
//   }
// }

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
  Circle  Circle   `xml:"circle"`
  Line    []Line     `xml:"line"`
}
type Circle  struct {
  Cx    float64 `xml:"cx,attr"`
  Cy    float64 `xml:"cy,attr"`
  R     float64 `xml:"r,attr"`
} 
type Line struct {
  X1    float64 `xml:"x1,attr"`
  Y1    float64 `xml:"y1,attr"`
  X2    float64 `xml:"x2,attr"`
  Y2    float64 `xml:"y2,attr"`
}

func TestSVGWriterAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
  b := bytes.Buffer{}

	clockface.SVGWriter(&b, tm)

  svg := SVG{}
  xml.Unmarshal(b.Bytes(), &svg)

  want := Line{150, 150, 150, 60}
  for _, line := range svg.Line{
    if line == want{
      return
    }
  }

  t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", want, svg.Line)
}

func TestSVGWriterSecondHand(t *testing.T){
  cases := []struct{
    time time.Time
    line Line
  }{
    {
      simpleTime(0, 0, 0),
      Line{150, 150, 150, 60},
    },
    {
      simpleTime(0, 0, 30),
      Line{150, 150, 150, 240},
    },
  }

  for _, c := range cases{
    t.Run(testName(c.time), func(t *testing.T){
      b := bytes.Buffer{}
      clockface.SVGWriter(&b, c.time)

      svg := SVG{}
      xml.Unmarshal(b.Bytes(), &svg)

      if !containsLine(c.line, svg.Line){
        t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
      }
    })
  }
}

func containsLine(l Line, ls []Line) bool{
  for _, line := range ls{
    if line == l{
      return true
    }
  }
  return false
}
