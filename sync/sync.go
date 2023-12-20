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

// Constructor which show readers of API that would be better to not initialise the type yourself
// bc of a property from mutex:
// -Mutex must not be copied after first use.
func NewCounter() *Counter{
  // return the addr of the instance of Counter
  return &Counter{}
}
