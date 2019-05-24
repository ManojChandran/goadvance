// Go channels
package main

import (
    "fmt"
    "time"
)

func main() {
  // channels are typed
  stringStream := make(chan string)
  go func ()  {
    stringStream <- "hello channels" // pass data into channel
  }()
  time.Sleep(5 * time.Millisecond)
  fmt.Println(<-stringStream) // receive data from channel
}
