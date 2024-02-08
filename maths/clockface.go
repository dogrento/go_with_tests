package clockface

import(
  "time"
)

type Point struct{
  X float64
  Y float64
}

func SecondHand(tm time.Time) Point{
  return Point{150, 60}
}

func secondHandPoint(t time.Time) Point{
  return Point{}
}
