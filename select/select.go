package main

import(
  "time"
  "net/http"
)

func Racer(a, b string) string{
  // select allow you to wait on multiple channels.
  // The first one to send a value "wins" and the code
  // underneath the case is executed.
  select{
  case <-ping(a):
    return a
  case <-ping(b):
    return b
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
