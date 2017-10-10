package main

import "fmt"
import "reflect"


var p = fmt.Println

func main() {
  var coba string = "gundar"
  // p(string(coba[0])) // ascii only

  p( string( []rune(coba)[1])  ) // utf8 only

  fmt.Println(reflect.TypeOf(coba[0]))

  fmt.Println(reflect.TypeOf([]rune(coba)[1]))

  p(coba[0])
  p(coba[1])
  p(coba[2])
  p(coba[3])
  p(coba[4])
  p(coba[5])
}
