package main

import "reflect"

func walk(x interface{}, fn func(input string)){
  // ValueOf returns a value of a given variable
  // so we can inspect a value
  val := reflect.ValueOf(x)
  // - and it's fields
  field := val.Field(0)

  fn(field.String())
}
