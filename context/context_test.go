package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
  "log"
)

type SpyStore struct{
  response  string
  t         *testing.T
}
func (s *SpyStore) Fetch(ctx context.Context) (string, error){
  data := make(chan string, 1)

  go func(){
    var result string
    
    //We are simulating a slow process where we build the result slowly by appending the string char by char in a goroutine
    // when goroutine finishes its work, it writes the string to the data chan
    // goroutine listens for the ctx.Doone and will stop the work if a signal is sent in that chan.
    for _, c := range s.response{
      select{
      case <-ctx.Done():
        log.Println("spy store go cancelled")
        return
      default:
        time.Sleep(10 * time.Millisecond)
        result += string(c)
      }
    }
    data <- result
  }()

  // Waits for goroutine to finish its job or for the cancellation to occur.
  select{
  case <-ctx.Done():
    return "", ctx.Err()
  case res := <-data:
    return res, nil
  }
}

func TestServer(t *testing.T){
  data := "hello world"

  t.Run("returns data from store", func(t *testing.T){
    store := &SpyStore{response: data, t: t}
    svr := Server(store)

    request := httptest.NewRequest(http.MethodGet, "/", nil)
    response := httptest.NewRecorder()

    svr.ServeHTTP(response, request)

    if response.Body.String() != data{
      t.Errorf("got %s, want %s", response.Body.String(), data)
    }
  })

  t.Run("tells store to cancel work if request is cancelled", func(t *testing.T){
    store := &SpyStore{response: data, t: t}  
    svr := Server(store)

    request := httptest.NewRequest(http.MethodGet, "/", nil)

    // Derive a new cancellingCtx from request which returns a cancel function
    cancellingCtx, cancel := context.WithCancel(request.Context())
    // Schedule the cancel function to be called in 5 ms
    time.AfterFunc(5 * time.Millisecond, cancel)
    // Using the new context in request by calling rq.withcontext
    request = request.WithContext(cancellingCtx)

    response := httptest.NewRecorder()

    svr.ServeHTTP(response, request)

    // if !store.cancelled{
    //   t.Error("store was not told to cancel")
    // }
    // store.assertWasCancelled()
  })
}
