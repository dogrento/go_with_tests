package main

import(
  "testing"
)

func TestSearch(t *testing.T){
  t.Run("searching key getting it's value", func(t *testing.T){
    // declaring map requires two types
    // - first is the key type which is written inside the []
    // - second is the value type, goes after []
    // NOTE: key type is a "comparable type"
    // dict := map[string]string{"test": "this is just a test"}
    dict := Dictionary{"test": "this is just a test"}

    got := dict.Search("test")
    want := "this is just a test"

    assertStrings(t, got, want)
  })
}

func assertStrings(t testing.TB, got, want string){
  t.Helper()

  if got != want{
    t.Errorf("\ngot -> %s\nwant -> %s", got, want)
  }

}
