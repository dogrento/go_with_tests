package main

func Sum(nums []int) int{
  var sum int
  for  s := range nums{
    sum += nums[s]
  }
  return sum 
}

func SumAll(slc1 []int, slc2 []int) []int{
  var response []int
  if slc2 != nil{
    response = append(response, Sum(slc1))
    response = append(response, Sum(slc2))
  } else{
    response = append(response, Sum(slc1))
  }
  return response
}
