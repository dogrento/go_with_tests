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

    defer slowServer.Close()
    defer fastServer.Close()

    slowUrl := slowServer.URL
    fastUrl := fastServer.URL

    want := fastUrl
    got, err := Racer(slowUrl, fastUrl)

    if err != nil{
      t.Fatalf("did not expected error> %v", err)
    }

    if got != want{
      t.Errorf("\ngot -> %s\nwant -> %s", got, want)
    }
  })

  t.Run("return err if a sever doesn't respond within 10s", func(t *testing.T){
    // serverA := makeDelayedServer(11 * time.Second)
    // serverB := makeDelayedServer(12 * time.Second)
    server := makeDelayedServer(25 * time.Millisecond)

    // defer will call the following function at the end of the containing function
    // -this benefits reading code bc it keeps the close operation (in this case),
    // close to its calling line
    // defer serverA.Close()
    // defer serverB.Close()
    defer server.Close()

    _, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

    if err == nil{
      t.Errorf("expected an error but didn't get one.")
    }
  })
}

func makeDelayedServer(delay time.Duration) *httptest.Server{
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    time.Sleep(delay)
    w.WriteHeader(http.StatusOK)
  }))
}
