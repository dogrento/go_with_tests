package main

import "testing"

func TestHello(t *testing.T){
  t.Run("saying hell to people", func(t *testing.T) {
    got := Hello("Moto!", "")
    want := "Hello, Moto!"

    assertCorrectMessage(t, got, want)
  })

  t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
    got := Hello("", "")
    want := "Hello, World"

    assertCorrectMessage(t, got, want)
  })

  t.Run("in Spanish", func(t *testing.T){
    got := Hello("Carmelita", "spanish")
    want := "Hola, Carmelita"
    assertCorrectMessage(t, got, want)
  })

  t.Run("in French", func(t *testing.T){
    got := Hello("Napoleon", "french")
    want := "Bonjour, Napoleon"
    assertCorrectMessage(t, got, want)
  })

  t.Run("in HUEHUE", func(t *testing.T){
    got := Hello("Bolsolula", "huehue")
    want := "Salve, Bolsolula"
    assertCorrectMessage(t, got, want)
  })
}

func assertCorrectMessage(t testing.TB, got, want string){
  // Helper method tells the test suite that this method is helper.
  // - when it fails the line number reported will be in out func call
  t.Helper()
  if got != want{
    t.Errorf("got: %q want: %q", got, want)
  }
}
