package main

import (
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
    got, _ := dict.Search("test")
    want := "this is just a test"

    assertStrings(t, got, want)
  })
  t.Run("Unknown key", func(t *testing.T){
    _, err := dict.Search("unknown")
    want := ErrKeyNotFound 

    if err == nil{
      t.Fatal("expected to get an error.")
    }

    // Error can be converted to a string with the Error()
    // assertStrings(t, err.Error(), want.Error())
    assertError(t, err, want)
  })
}

func TestInsert(t *testing.T){
  dict := Dictionary{"test": "this is just a test"}

  t.Run("Insert key", func(t *testing.T){
    key := "otherKey"
    value := "otherValue"

    dict.Insert(key, value)

    assertValue(t, dict, key, value)
  })
  
  t.Run("Inserting existing key", func(t *testing.T){
    err := dict.Insert("test", "this is just a test")

    assertError(t, err, ErrWordExists)
    assertValue(t, dict, "test", "this is just a test")
  })
}

func TestUpdate(t *testing.T){
  key := "test1"
  value := "test1_value"
  dict := Dictionary{key: value}

  t.Run("Updating dict", func(t *testing.T){
    newValue := "new_test_value"

    dict.Update(key, newValue)

    assertValue(t, dict, key, newValue)
  })

  t.Run("key that does not exist", func(t *testing.T){
    otherKey := "test2"
    otherValue := "test2_value"

    err := dict.Update(otherKey, otherValue)

    assertError(t, err, ErrWordDoesNotExists)
  })
}

func TestDelete(t *testing.T){
  t.Run("Delete key from dict", func(t *testing.T){
    key := "test"
    dict := Dictionary{key: "test_value"}

    dict.Delete(key)

    _, err := dict .Search(key)
    if err != ErrKeyNotFound{
      t.Errorf("Expected %q to be deleted", key)
    }
  })
}

func assertStrings(t testing.TB, got, want string){
  t.Helper()

  if got != want{
    t.Errorf("\ngot -> %s\nwant -> %s", got, want)
  }
}

func assertValue(t testing.TB, dict Dictionary, key, value string){
  t.Helper()

  got, err := dict.Search(key)
  if err != nil{
    t.Fatalf("Expected to find: %q. %q", key, err)
  }
  assertStrings(t, got, value)
}

func assertError(t testing.TB, got, want error){
  t.Helper()

  if got != want{
    t.Errorf("\ngot error -> %q\nwant error -> %q", got, want)
  }
}
