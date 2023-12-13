package main

// creating a wrapper Dictionary around map 
type Dictionary map[string]string

func (d Dictionary) Search(key string) string{
  return d[key] 
}
