package main

import (
	"bytes"
	"reflect"
	"testing"
  "time"
)

func TestCountdown(t *testing.T){
  t.Run("testing Countdown()", func (t *testing.T){
    buffer := &bytes.Buffer{}
    SpySleeper := &SpySleeper{}

    Countdown(buffer, SpySleeper)

    got := buffer.String()
    // backtick is another way of creating a string but lets you include things like newlines
    want := `3
2
1
Go!`

    if got != want{
      t.Errorf("\ngot->%s\nwant->%s", got, want)
    }

    if SpySleeper.Calls != 3{
      t.Errorf("Not enough calls to sleeper.\ngot->%d\nwant->3", SpySleeper.Calls)
    }
  })

  t.Run("testing Calls", func(t *testing.T){
    spySleepPrinter := &SpyCountdownOperations{}
    Countdown(spySleepPrinter, spySleepPrinter)

    want := []string{
      write,
      sleep,
      write,
      sleep,
      write,
      sleep,
      write,
    }

    if !reflect.DeepEqual(want, spySleepPrinter.Calls){
      t.Errorf("\nwanted calls %v\ngot %v", want, spySleepPrinter.Calls)
    }
  })
}

func TestConfigurableSleeper(t *testing.T){
  t.Run("testing configurable sleeper", func(t *testing.T){
    sleepTime := 5 * time.Second

    spyTime := &SpyTime{}
    sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
    sleeper.Sleep()

    if spyTime.durationSlept != sleepTime{
      t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
    }
  })
}
