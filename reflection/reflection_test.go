package main

import (
	"reflect"
	"testing"
)

type Person struct{
  Name    string
  Profile Profile
}

type Profile struct{
  Age  int
  City string
}

func TestWalk(t *testing.T){

  cases := []struct{
    Name          string
    Input         interface{}
    ExpectedCalls []string
  }{
    {
      "struct with one string field",
      struct{
        Name string
      }{"dgt"},
      []string{"dgt"},
    },
    {
      "struct with two string fields",
      struct{
        Name string
        City string
      }{"dgt", "curitiba"},
      []string{"dgt", "curitiba"},
    },
    {
      "struct with non string field",
      struct{
        Name string
        Age int
      }{"dgt", 30},
      []string{"dgt"},
    },
    {
      "struct with nested fields",
      Person{
        "dgt",
        Profile{30, "curitiba"},
      },
      []string{"dgt", "curitiba"},
    },
    {
      "pointers to things",
      &Person{
        "dgt",
        Profile{30, "curitiba"},
      },
      []string{"dgt", "curitiba"},
    },
    {
      "slices",
      []Profile{
        {30, "curitiba"},
        {33, "sao paulo"},
      },
      []string{"curitiba", "sao paulo"},
    },
    {
      "arrays",
      [2]Profile{
        {30, "curitiba"},
        {33, "sao paulo"},
      },
      []string{"curitiba", "sao paulo"},
    },
  }

  for _, test := range cases{
    t.Run(test.Name, func(t *testing.T){
      var got []string
      walk(test.Input, func(input string){
        got = append(got, input)
      })

      if !reflect.DeepEqual(got, test.ExpectedCalls){
        t.Errorf("\ngot -> %v\nwant -> %v", got, test.ExpectedCalls)  
      }
    })
  }

  t.Run("Maps", func(t *testing.T){
    aMap := map[string]string{
      "cow": "moo",
      "sheep": "baa",
    }

    var got []string
    walk(aMap, func(input string){
      got = append(got, input)
    })

    assertContains(t, got, "moo")
    assertContains(t, got, "baa")
  })
  
  t.Run("Channels", func(t *testing.T){
    aChannel := make(chan Profile)

    go func(){
      aChannel <- Profile{30, "curitiba"}
      aChannel <- Profile{34, "sao paulo"}
      close(aChannel)
    }()

    var got []string
    want := []string{"curitiba", "sao paulo"}

    walk(aChannel, func(input string){
      got = append(got, input)
    })

    if !reflect.DeepEqual(got, want){
      t.Errorf("got %v, want %v", got, want)
    }
  })
  
  t.Run("Func", func(t *testing.T){
    aFunction := func() (Profile, Profile){
      return Profile{30, "curitiba"}, Profile{34, "sao paulo"}
    }

    var got []string
    want := []string{"curitiba", "sao paulo"}

    walk(aFunction, func(input string){
      got = append(got, input)
    })

    if !reflect.DeepEqual(got, want){
      t.Errorf("got %v, want %v", got, want)
    }
  })
}

func assertContains(t testing.TB, got []string, str string){
  t.Helper()
  contains := false
  for _, x := range got{
    if x == str{
      contains = true
    }
  }

  if !contains{
    t.Errorf("expected %v to contain %q but it didnt", got, str)
  }
}
