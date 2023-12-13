package main

import (
	"errors"
)

var ErrKeyNotFound = errors.New("Key not found!")
// creating a wrapper Dictionary around map 
type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error){
  value, ok := d[key]
  if !ok{
    return "", ErrKeyNotFound
  }
  return value, nil
}
