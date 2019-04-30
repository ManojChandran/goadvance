// Greatest Common devisor
package main

import (
  "fmt"
)

func main(){
  fmt.Println(gcd(3,6))
  fmt.Println(gcd(2,3))
}

func gcd(x, y int) int {
  for y != 0 {
    x,y = y, x%y
  }
  return x
}
