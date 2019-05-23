// Go routine example
package main

import (
  "fmt"
  "time"
)

func main()  {
  var msg = "Hello"
  go func(){
    fmt.Println(msg) // Prints "good bye"
  }()
  msg = "good bye"
  //sleep time introduced for go routine to complete
  time.Sleep(100 * time.Millisecond)
}
