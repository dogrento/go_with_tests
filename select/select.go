package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error){
  return ConfigurableRacer(a, b, tenSecondTimeout)
}
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error){
  // select allow you to wait on multiple channels.
  // The first one to send a value "wins" and the code
  // underneath the case is executed.
  select{
  case <-ping(a):
    return a, nil
  case <-ping(b):
    return b, nil
  case <-time.After(timeout):
    return "", fmt.Errorf("timed out wainting for %s and %s", a, b)
  }
}

func ping(url string) chan struct{} {
  ch := make(chan struct{})
  go func(){
    http.Get(url)
    close(ch)
  }()
  return ch
}

func measureResponseTime(url string) time.Duration{
  start := time.Now()
  http.Get(url)
  
  return time.Since(start)
}
