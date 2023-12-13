package main

import (
	"errors"
)

var ErrKeyNotFound = errors.New("Key not found!")
// creating a wrapper Dictionary around map 
type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error){
  // map lookup returns 2 values
  // - first: the value itself, if none then the value= ""
  // - second: a boolean
  value, ok := d[key]
  if !ok{
    return "", ErrKeyNotFound
  }
  return value, nil
}

func (d Dictionary) Insert(key, value string){
  d[key] = value
}
