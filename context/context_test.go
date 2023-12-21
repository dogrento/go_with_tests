package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct{
  response  string
  cancelled bool
  t         *testing.T
}
func (s *SpyStore) assertWasCancelled(){
  s.t.Helper()
  if!s.cancelled{
    s.t.Error("store was not told to cancel")
  }
}
func (s *SpyStore) assertWasNotCancelled(){
  s.t.Helper()
  if s.cancelled{
    s.t.Error("store was told to cancel")
  }
}
func (s *SpyStore) Fetch() string{
  time.Sleep(100 * time.Millisecond)
  return s.response
}
func (s *SpyStore) Cancel(){
  s.cancelled = true
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

    // if store.cancelled{
    //   t.Error("it should not have cancelled the store")
    // }
    store.assertWasNotCancelled()
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
    store.assertWasCancelled()

  })
}
