package main

import "reflect"

func walk(x interface{}, fn func(input string)){
  // ValueOf returns a value of a given variable
  // so we can inspect a value
  // - and it's fields
  val := reflect.ValueOf(x)

  for i := 0; i < val.NumField(); i++{
    field := val.Field(i)

    if field.Kind() == reflect.String{
      fn(field.String())
    }
  }
}
