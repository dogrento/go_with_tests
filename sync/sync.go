package main

import "sync"

type Counter struct{
  mu    sync.Mutex
  value int
}

func(c *Counter) Inc(){
  // Any goroutine calling this func, will acquire the lock on Counter if they're first
  // all the other goroutines will have to wait for it to be Unlocked before getting access
  c.mu.Lock()
  defer c.mu.Unlock()
  c.value++
}
func(c *Counter) Value() int{
  return c.value
}
