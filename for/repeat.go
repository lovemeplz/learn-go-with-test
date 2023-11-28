package iteration

import "fmt"

func Repeat(character string, times int) string {
  var repeated string
  for i := 0; i < times; i++ {
    repeated += character
  }
  return repeated
}

func ExampleRepeat() {
  repeat := Repeat("a", 5)
  fmt.Println(repeat)
}
