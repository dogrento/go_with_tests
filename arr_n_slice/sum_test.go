package main

import "testing"

func TestSum(t *testing.T){
  num := [5]int{1, 2, 3, 4, 5}

  got := Sum(num)
  want := 15

  if got != want{
    t.Errorf("\ngot -> %d \nwant -> %d\ngiven -> %v", got, want, num)
  }
}
