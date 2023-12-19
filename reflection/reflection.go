package main

import "reflect"

func walk(x interface{}, fn func(input string)){
  val := getValue(x)

  numberOfValues := 0
  var getField func(int) reflect.Value

  switch val.Kind(){
  // if it's a struct or a slice, we iterate over its values
  case reflect.Slice:
  // for i := 0; i < val.Len(); i++{
  //   walk(val.Index(i).Interface(), fn)
  // }
    numberOfValues = val.Len()
    getField = val.Index
  case reflect.Struct:
  // for i := 0; i < val.NumField(); i++{
  //   walk(val.Field(i).Interface(), fn)
  // }
    numberOfValues = val.NumField()
    getField = val.Field
  case reflect.String:
    fn(val.String())
  }

  for i := 0; i < numberOfValues; i++{
    walk(getField(i).Interface(), fn)
  }
}

func getValue(x interface{}) reflect.Value{
  // ValueOf returns a value of a given variable
  // so we can inspect a value
  // - and it's fields
  val := reflect.ValueOf(x)     

  // cant use NumField on a ptr value
  // it is needed to extract the underlying value before we can do that with Elem()
  if val.Kind() == reflect.Pointer{
    val = val.Elem()
  }

  return val
}
