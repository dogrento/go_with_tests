package main

import "reflect"

func walk(x interface{}, fn func(input string)){
  val := getValue(x)
  // DRYs up the calls to walk()
  // they only have to extract out the reflect.Value from val
  walkValue := func(value reflect.Value){
    walk(value.Interface(), fn)
  }

  switch val.Kind(){
  // if it's a struct or a slice, we iterate over its values
  case reflect.Slice, reflect.Array:
    for i := 0; i < val.Len(); i++{
      walkValue(val.Index(i))
    }
  case reflect.Struct:
    for i := 0; i < val.NumField(); i++{
      walkValue(val.Field(i))
    }
  case reflect.String:
    fn(val.String())
  case reflect.Map:
    for _, key := range val.MapKeys(){
      walkValue(val.MapIndex(key))
    }
  case reflect.Chan:
    for v, ok := val.Recv(); ok; v, ok = val.Recv(){
      walkValue(v)
    }
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
