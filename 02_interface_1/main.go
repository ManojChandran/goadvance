package main

import (
    "fmt"
)

type Reader interface {
  Read(p []byte) (n int, err error)
}

func main() {
    fmt.Println("hello")
}
