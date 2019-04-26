package main

import (
  "fmt"
)

func main(){

  var x interface{}
  x = 3.14
  fmt.Printf("x: type = %T, value = %v\n",x,x)
  goo := x
  fmt.Printf("x: type = %T, value = %v\n",goo,goo)
}
