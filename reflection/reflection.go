package main

import "reflect"

func walk(x interface{}, fn func(input string)){
  val := getValue(x)

  switch val.Kind(){
  case reflect.Slice:
  for i := 0; i < val.Len(); i++{
    walk(val.Index(i).Interface(), fn)
  }
  case reflect.Struct:
  for i := 0; i < val.NumField(); i++{
    walk(val.Field(i).Interface(), fn)
  }
  case reflect.String:
    fn(val.String())
  }
  // if val.Kind() == reflect.Slice{
  //   for i := 0; i < val.Len(); i++{
  //     walk(val.Index(i).Interface(), fn)
  //   }

  //   return
  // }

  // for i := 0; i < val.NumField(); i++{
  //   field := val.Field(i)

  //   switch field.Kind(){
  //   case reflect.String:
  //     fn(field.String())
  //   case reflect.Struct:
  //     walk(field.Interface(), fn)
  //   }
  // }
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
