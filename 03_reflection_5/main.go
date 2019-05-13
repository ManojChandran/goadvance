package main

import (
  "fmt"
  "reflect"
)

func main() {
  var x interface{}
  x = 3.14

  t0 := reflect.TypeOf(x)
  v0 := reflect.TypeOf(x)

  fmt.Printf("x: type = %v, value = %v\n", t0, v0)
  goo := x
  fmt.Printf("x: type = %v, value = %v\n", goo, goo)

  x = &struct{ name string }{}
  t1 := reflect.TypeOf(x)
  v1 := reflect.TypeOf(x)
  fmt.Printf("x: type = %v, value = %v\n", t1, v1)
  hoo := x
  fmt.Printf("x: type = %T, value = %v\n", hoo, hoo)

  // what kind or category is the type a member of?
  //
  fmt.Printf("x: type = %T, value = %v\n", t0,t0.Kind())
  fmt.Printf("x: type = %T, value = %v\n", t1, t1.Kind())

}
