/*
Preventing goroutine leaks
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
  dowork := func (done <-chan interface{}, strings <-chan string) <-chan interface{} {
    completed := make(chan interface{})
    go func() {
      defer fmt.Println("done exit")
      defer close(completed)
      fmt.Println("in routine")
      for {
        select{
        case s:= <-strings:
          fmt.Println(s)
        case <-done:
            return
        }
      }
    }()
    return completed
  }
  done := make(chan interface{})
  dowork(done, nil)
  go func ()  {
    time.Sleep(1*time.Second)
    fmt.Println("close dowork goroutine")
    close(done)
  }()
  time.Sleep(5*time.Second)
  fmt.Println("done")
}
