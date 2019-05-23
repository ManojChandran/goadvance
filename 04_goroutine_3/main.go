// Go Mutex example
package main

import (
  "fmt"
  "sync"
)

var count int
var lock sync.Mutex
var arithmatic sync.WaitGroup

func increment()  {
  lock.Lock()
  defer lock.Unlock()
  count++
  fmt.Printf("Incrementing : %d\n", count)
}

func decrement()  {
  lock.Lock()
  defer lock.Unlock()
  count--
  fmt.Printf("Decrementing : %d\n", count)
}

func main()  {
  for i:=0; i <= 5; i++{
      arithmatic.Add(1)
      go func ()  {
        defer arithmatic.Done()
        increment()
      }()
  }

  for i:=0; i <= 5; i++{
      arithmatic.Add(1)
      go func ()  {
        defer arithmatic.Done()
        decrement()
      }()
  }
  arithmatic.Wait()
  fmt.Println("Done...")
}
