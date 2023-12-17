package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T){
  t.Run("racer test", func(t *testing.T){
    slowServer := makeDelayedServer(20 * time.Millisecond)
    fastServer := makeDelayedServer(0 * time.Millisecond)

    slowUrl := slowServer.URL
    fastUrl := fastServer.URL

    want := fastUrl
    got := Racer(slowUrl, fastUrl)

    if got != want{
      t.Errorf("\ngot -> %s\nwant -> %s", got, want)
    }

    slowServer.Close()
    fastServer.Close()
  })
}

func makeDelayedServer(delay time.Duration) *httptest.Server{
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    time.Sleep(delay)
    w.WriteHeader(http.StatusOK)
  }))
}
