package main

import "fmt"

func Sum(nums [5]int) int{
  var sum int
  for  s := range nums{
    sum += nums[s]
  }
  return sum 
}

func main(){
  var arr = [5]int {1, 2, 3, 4, 5}
  fmt.Println(Sum(arr))
}
