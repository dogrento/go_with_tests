package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T){
  // t.Run("testing Walk()", func (t *testing.T){

  //   expected := "dgt"
  //   var got []string

  //   // creating anonymous struct
  //   x := struct{
  //     Name string
  //   }{expected}

  //   // walk with spy
  //   walk(x, func(input string){
  //     got = append(got, input)
  //   })
  //   
  //   if len(got) != 1{
  //     t.Errorf("wrong number of function calls\ngot -> %d\nwant -> %d", len(got), 1)
  //   }
  //   if got[0] != expected{
  //     t.Errorf("\ngot -> %q\nwant -> %q", got[0], expected)
  //   }
  // })

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
