// go routine pitfall 2
package main

import (
  "fmt"
  "time"
)

func main() {
  stringStream := make (chan string)
  go func() {
//  closing the channel might be a best option
//  we can read from closed channels    
//    defer close(stringStream)
    if 0 != 1{
      return
    }
    stringStream <- "hello channel"
  }()
  time.Sleep(100 * time.Millisecond)
  fmt.Println(<-stringStream)
}

/*
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/home/manoj/github.com/ManojChandran/goadvance/05_goroutine_pitfall_2/main.go:18 +0x7c
exit status 2
*/

/*
Channel never receive's value because of the condition.
When anonymous goroutine exits, GO detects that all goroutines
are asleep and reports a deadlock
*/
