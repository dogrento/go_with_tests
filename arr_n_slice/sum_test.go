package main

import "testing"

func TestSum(t *testing.T){
  t.Run("collection of any size", func(t *testing.T){
    numbers := []int{1, 2, 3}

    got := Sum(numbers)
    want := 6

    if got != want{
      t.Errorf("\ngot -> %d \nwant -> %d\ngiven -> %v", got, want, numbers)
    }
  })
}

func TestSumAll(t *testing.T){
  arr1 := []int{1, 2}
  arr2 := []int{0, 9}
  arr3 := []int{1, 1, 1}

  got1 := SumAll(arr1, arr2)
  want1 := []int{3, 9} 

  got2 := SumAll(arr3, nil)
  want2 := []int{3}

  assertMsg(t, got1, want1, arr1, arr2)
  assertMsg(t, got2, want2, arr3, nil)
}
  
func assertMsg(t testing.TB, got []int, want []int, num1 []int, num2 []int){
  t.Helper()
  if num2 != nil{
    if isSlicesEqual(got, want) == true{
      t.Errorf("\ngot -> %d \nwant -> %d\ngiven -> %v\n%v", got, want, num1, num2)
    }
  }
}

func isSlicesEqual(arr1 []int, arr2 []int) bool{
  if len(arr2) != len(arr2){
    return false
  }

  for i, v := range arr1{
    if v != arr2[i]{
      return false
    }
  }
  return true
}
