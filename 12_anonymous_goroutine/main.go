package main

import (
  "fmt"
  "runtime"
)

func main(){
  fmt.Println("Outside a go routine")   // declares anonymous
  go func ()  {                         // function and executes
    fmt.Println("Inside a go routine")  // as a go routine
  }()

  fmt.Println("Outside again")
  runtime.Gosched()
}
