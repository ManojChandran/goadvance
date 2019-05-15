package main

import (
    "./encapsulation"
)

func main() {
  e := encapsulation.Encapsulation{}

  e.Expose()

  e.Unhide()
}
