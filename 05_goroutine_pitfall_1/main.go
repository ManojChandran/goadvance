// go routine pitfall 1
/*
  go has both closure and go routines
  closure variable are evaluated when the goroutine is run
*/
package main

import (
  "fmt"
  "time"
)

func main() {
  for i:=0; i<=9; i++{
    go func(){
        fmt.Println(i) // prints 10
    }()
  }
  time.Sleep(100 * time.Millisecond)
}
