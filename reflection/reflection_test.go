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
}
