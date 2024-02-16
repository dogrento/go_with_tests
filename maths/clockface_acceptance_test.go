package clockface

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"
)

// SVG was generated 2024-02-16 16:37:18 by https://xml-to-go.github.io/ in Ukraine.
type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  struct {
		Text  string `xml:",chardata"`
		Cx    string `xml:"cx,attr"`
		Cy    string `xml:"cy,attr"`
		R     string `xml:"r,attr"`
		Style string `xml:"style,attr"`
	} `xml:"circle"`
	Line struct {
		Text  string `xml:",chardata"`
		X1    string `xml:"x1,attr"`
		Y1    string `xml:"y1,attr"`
		X2    string `xml:"x2,attr"`
		Y2    string `xml:"y2,attr"`
		Style string `xml:"style,attr"`
	} `xml:"line"`
} 

func TestSVGWriterAtMidnight(t *testing.T){
  tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

  b := bytes.Buffer{}
  clockface.SVGWriter(&b, tm)

  svg := SVG{}
  xml.Unmarshal(b.Bytes(), &svg)

  x2 := "150"
  y2 := "60"

  for _, line := range svg.Line{
    if line.X2 == x2 && line.Y2 == y2{
      return
    }
  }

  t.Errorf("Expected to find the second hand with x2 of %+v and y2 of %+v, in the SVG output %v", x2, y2, b.String())
}

func TestSecondsInRadians(t *testing.T){
  cases := []struct{
    time time.Time
    angle float64
  }{
    {simpleTime(0, 0, 30), math.Pi},
    {simpleTime(0, 0, 0), 0},
    {simpleTime(0, 0, 45), (math.Pi / 2) * 3},
    {simpleTime(0, 0, 7), (math.Pi / 30) * 7},
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
  return (math.Pi / (30 / (float64(tm.Second()))))
}

func simpleTime(h, min, sec int) time.Time{
  return time.Date(312, time.October, 28, h, min, sec, 0, time.UTC)
}

func testName(t time.Time) string{
  return t.Format("15:04:05")
}
