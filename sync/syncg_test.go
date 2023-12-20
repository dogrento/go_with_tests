package main

import (
	"sync"
	"testing"
)

func TestCounter(t * testing.T){
  t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T){
    // counter := Counter{}
    counter := NewCounter()
    counter.Inc()
    counter.Inc()
    counter.Inc()

    assertCounter(t, counter, 3)
  })
  t.Run("runs safely concurrently", func(t *testing.T){
    wantedCount := 1000
    // counter := Counter{}
    counter := NewCounter()

    // waitgroup waits for a collection of goroutines to finish
    var wg sync.WaitGroup
    // calls Add to set the number of goroutines to wait for
    wg.Add(wantedCount)

    for i := 0; i < wantedCount; i++{
      go func(){
        counter.Inc()
        // each of goroutines runs and calls Done when finished
        wg.Done()
      }()
    }

    // Wait() can be used to block until all goroutines have finished
    // we can be sure all of our goroutines have attempted to Inc the Counter
    wg.Wait()

    assertCounter(t, counter, wantedCount)
  })
}

func assertCounter(t testing.TB, got *Counter, want int){
  t.Helper()
  if got.Value() != want{
    t.Errorf("got %d, want %d", got.Value(), want)
  }

}
