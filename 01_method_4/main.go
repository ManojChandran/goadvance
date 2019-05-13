package main

import (
    "fmt"
    "math"
)

type Point struct {
  X,Y float64
}
func (p Point) Abs() float64  {
  return math.Sqrt(p.X * p.X + p.Y * p.Y)
}

func main() {
    var p *Point
    fmt.Println(p.Abs()) // invalid memory address or nil pointer dereference
}
/*
Un initialized pointer pin the main function is nil, and you can't follow
the nil pointer.
if x is nil, an attempt to evaluate *x will cause a run-time panic
*/
