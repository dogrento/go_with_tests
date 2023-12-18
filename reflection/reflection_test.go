package main

import(
  "testing"
)

func TestWalk(t *testing.T){
  t.Run("testing Walk()", func (t *testing.T){

    expected := "dgt"
    var got []string
    x := struct{
      Name string
    }{expected}

    walk(x, func(input string){
      got = append(got, input)
    })
    
    if len(got) != 1{
      t.Errorf("wrong number of function calls\ngot -> %d\nwant -> %d")
    }
  })
}
