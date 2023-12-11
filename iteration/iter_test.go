package iteration

import (
	"fmt"
	"testing"
)


func TestRepeat(t *testing.T){
  value := 5
  repeated := Repeat("a", value)
  expected := "aaaaa"

  if repeated != expected{
    t.Errorf("expected -> %q ; got -> %q", expected, repeated)
  }
}

// to run test: go test -v
func ExampleRepeat(){
  example := Repeat("a", 5)
  fmt.Println(example)
  // Output: aaaaa
}

// to run benchmark: go test -bench="."
func BenchmarkRepeat(b *testing.B){
  // where b.N is determined by benchmark system
  for i := 0; i< b.N; i++ {
    Repeat("a", 5)
  }
}
