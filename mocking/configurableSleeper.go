package main

import "time"

type ConfigurableSleeper struct{
  duration time.Duration
  sleep    func(time.Duration)  
}

func (c *ConfigurableSleeper) Sleep(){
  c.sleep(c.duration)
}
