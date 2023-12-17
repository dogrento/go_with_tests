package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T){
  t.Run("racer test", func(t *testing.T){
    slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
      time.Sleep(20 * time.Millisecond)
      w.WriteHeader(http.StatusOK)
    }))

    fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, t *http.Request){
      w.WriteHeader(http.StatusOK)
    }))

    // slowUrl := "http://www.facebook.com"
    // fastUrl := "http://www.quii.dev"
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
