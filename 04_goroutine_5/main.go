package main

import (
  "fmt"
  "runtime"
)

func main()  {
  fmt.Println(" Main function ")
  go func ()  {
    fmt.Println(" in goroutine ")
  }()
  fmt.Println(" back in Main function ")

  runtime.Gosched() // gives an opportunity
                    //to execute other goroutines 
}
