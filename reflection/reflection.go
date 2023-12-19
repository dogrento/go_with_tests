package main

import "reflect"

func walk(x interface{}, fn func(input string)){
  // ValueOf returns a value of a given variable
  // so we can inspect a value
  // - and it's fields
  val := reflect.ValueOf(x)

  if val.Kind() == reflect.Pointer{
    val = val.Elem()
  }

  for i := 0; i < val.NumField(); i++{
    field := val.Field(i)

    switch field.Kind(){
    case reflect.String:
      fn(field.String())
    case reflect.Struct:
      walk(field.Interface(), fn)
  }
  }
}
