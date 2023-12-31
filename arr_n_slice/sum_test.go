package main

import (
	"reflect"
	"testing"
)

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

  t.Run("Running with assrtMsg()", func(t *testing.T){
    assertMsg(t, got1, want1, arr1, arr2)
    assertMsg(t, got2, want2, arr3, nil)
  })

  t.Run("Runnig with reflect.DeepEqual()", func(t *testing.T){
    if !reflect.DeepEqual(got1, want1){
      t.Errorf("got -> %v\nwant-> %v", got1, want1)
    }

    if !reflect.DeepEqual(got2, want2){
      t.Errorf("got -> %v\nwant-> %v", got2, want2)
    }
  })
}

func TestSumAllTails(t *testing.T){
  // Assign function to a var
  // - It can be useful to bind a func to other local var in scope
  // - Also allos to reduce the surface area of API
  // - It can not be accessed by other dunc in this package
  checkSums := func(t testing.TB, got, want []int){
    t.Helper()

    if !reflect.DeepEqual(got, want){
      t.Errorf("got -> %v \nwant -> %v", got, want)
    }
  }

  t.Run("Sum tails", func(t *testing.T){
    arr1 := []int{1, 2}
    arr2 := []int{3, 4, 1}

    got := SumAllTails(arr1, arr2)
    want := []int{2, 5}

    checkSums(t, got, want)
  })

  t.Run("Sum tails from 3 slices", func(t *testing.T){
    arr1 := []int{1, 1, 1}
    arr2 := []int{1, 1, 1}
    arr3 := []int{1, 1, 1}

    got := SumAllTails(arr1, arr2, arr3)
    want := []int{2, 2, 2}

    checkSums(t, got, want)
  })

  t.Run("Sum tails empty slices", func(t *testing.T){
    got := SumAllTails([]int{}, []int{3, 4, 5})
    want := []int{0, 9}
    if !reflect.DeepEqual(got, want){
      t.Errorf("got -> %v \nwant -> %v", got, want)
    }
  })
}

func TestIsSliceEqual(t *testing.T){
  t.Run("Getting a true validation", func(t *testing.T){
    arr1 := []int{1, 2, 3}
    arr2 := []int{1, 2, 3}

    got := isSlicesEqual(arr1, arr2)
    want := true

    if got != want{
      t.Errorf("got -> %t\n want -> %t", got, want)
    }
  })
}
  
func assertMsg(t testing.TB, got []int, want []int, num1 []int, num2 []int){
  t.Helper()
  if num2 != nil{
    if !isSlicesEqual(got, want){
      t.Errorf("\ngot -> %d \nwant -> %d\ngiven -> %v\n%v", got, want, num1, num2)
    }
  } else {
    if !isSlicesEqual(got, want){
      t.Errorf("\ngot -> %d \nwant -> %d\ngiven -> %v\n", got, want, num1)
    }
  }
}

func isSlicesEqual(arr1 []int, arr2 []int) bool{
  if len(arr1) != len(arr2){
    return false
  }

  for i, v := range arr1{
    if v != arr2[i]{
      return false
    }
  }
  return true
}
