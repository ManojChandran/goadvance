// Go routine example
package main

import (
  "fmt"
  "sync"
)

var wg sync.WaitGroup

func main() {
  msg := "Hello"
  wg.Add(1) // add the below go routine to WaitGroup
  go func(){
      defer wg.Done() // reports WaitGroup when function done
      fmt.Println(msg)
  }()
  msg = "Good bye"
  wg.Wait() // wait for the routine to complete
}
