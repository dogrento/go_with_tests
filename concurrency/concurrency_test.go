package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool{
  if url == "waat://furhurterwe.geds"{
    return false
  }

  return true
}

func slowWebsiteChecker(_ string) bool{
  time.Sleep(20 * time.Millisecond)
  return true
}

func BenchmarkCheckWebsites(b *testing.B){
  // creating a slice of string type with 100 elements
  urls := make([]string, 100)
  for i := 0; i < b.N; i++{
    CheckWebsites(slowWebsiteChecker, urls)
  }
}

func TestCheckWebsites(t *testing.T){
  t.Run("testing website checker", func(t *testing.T){
    websites := []string{
      "http://google.com",
      "http://blog.gypsydave5.com",
      "waat://furhurterwe.geds",
    }

    want := map[string]bool{
      "http://google.com":          true,
      "http://blog.gypsydave5.com": true,
      "waat://furhurterwe.geds":    false,
    }

    got := CheckWebsites(mockWebsiteChecker, websites)

    if !reflect.DeepEqual(want, got){
      t.Fatalf("\nwant -> %v\ngot -> %v", want, got)
    }
  })
}
