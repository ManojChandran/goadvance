/*
Goroutine leaks
*/
package main

import (
  "fmt"
  "time"
)

func main()  {
  fmt.Println("hello")
  // main go routine passes nil channel into dowork
  // strings channel never get anything written to it
  // goroutine remains in memmory for the life time
  // of this proess
  dowork := func (strings <-chan string) <-chan interface{} {
    completed := make(chan interface{})
    go func ()  {
      defer fmt.Println("done exit")
      defer close(completed)
      fmt.Println("in routine")
      for s:= range strings {
        fmt.Println(s)
      }
    }()
    return completed
  }

  dowork(nil)
  time.Sleep(5*time.Second)
  fmt.Println("done")
}
