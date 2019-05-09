// go methods
/*
  Receiver can be pointer
*/
package main

import (
    "fmt"
)

type rect struct {
  width, height float64
}

func (r *rect) perim() float64 {
  return 2 * r.width + 2*r.height
}

func main() {
  r := rect{width: 10, height: 5}
  perimeter := r.perim()
  fmt.Println("Perimeter :", perimeter)
}
