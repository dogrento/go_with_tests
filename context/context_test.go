package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyStore struct{
  response  string
  cancelled bool
}
func (s *SpyStore) Fetch() string{
  return s.response
}
func (s *SpyStore) Cancel(){
  s.cancelled = true
}

func TestServer(t *testing.T){
  t.Run("testing server response", func(t *testing.T){
    data := "hello world"
    svr := Server(&SpyStore{data})

    request := httptest.NewRequest(http.MethodGet, "/", nil)
    response := httptest.NewRecorder()

    svr.ServeHTTP(response, request)

    if response.Body.String() != data{
      t.Errorf("got %s, want %s", response.Body.String(), data)
    }
  })
}
