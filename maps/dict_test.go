package main

import(
  "testing"
)

func TestSearch(t *testing.T){
  // declaring map requires two types
  // - first is the key type which is written inside the []
  // - second is the value type, goes after []
  // NOTE: key type is a "comparable type"
  // dict := map[string]string{"test": "this is just a test"}
  dict := Dictionary{"test": "this is just a test"}

  t.Run("Known key", func(t *testing.T){
    got := dict.Search("test")
    want := "this is just a test"

    assertStrings(t, got, want)
  })
   t.Run("Unknown key", func(t *testing.T){
    _, err := dict.Search("unknown")
    want := "could not find the word you were looking for"

    if err == nil{
      t.Fatal("expected to get an error.")
    }
  })
}

func assertStrings(t testing.TB, got, want string){
  t.Helper()

  if got != want{
    t.Errorf("\ngot -> %s\nwant -> %s", got, want)
  }

}
