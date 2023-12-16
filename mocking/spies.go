package main

import(
  "time"
)

//MOCK
// Spies are a kind of mock which can record how a dependency is used
// they can record the arguments sent in, how many times it has been called, etc.
// In this case, this gonna track of how many time Sleep() is called
type SpySleeper struct{
  Calls int
}
func (s *SpySleeper) Sleep(){
  s.Calls++
}

const (
  write = "write"
  sleep = "sleep"
)
type SpyCountdownOperations struct{
  Calls []string
}
func (s *SpyCountdownOperations) Sleep(){
  s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error){
  s.Calls = append(s.Calls, write)
  return
}

type SpyTime struct{
  durationSlept time.Duration
}
func (s *SpyTime) Sleep(duration time.Duration){
  s.durationSlept = duration
}


