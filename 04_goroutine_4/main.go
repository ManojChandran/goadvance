// GOMAXPROCS
/*
  function controls the number of OS
  threads that will host so-called
  work queues
*/
package main

import (
  "fmt"
  "runtime"
)

func main()  {
  fmt.Printf("Threads: %v\n",runtime.GOMAXPROCS(runtime.NumCPU()))
  fmt.Printf("Threads: %v\n",runtime.GOMAXPROCS(-1))
  runtime.GOMAXPROCS(1) // forcing to have single core
  fmt.Printf("Threads: %v\n",runtime.GOMAXPROCS(-1))
}
