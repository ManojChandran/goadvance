// go channel
package main

import (
  "fmt"
)

func main(){
  inStream := make(chan int)
  go func(){
      defer close(inStream)
      for i:=1; i<=5; i++{
        inStream<- i
      }
  }()
  // range over the channel
  for integer := range inStream {
    fmt.Printf("%v\n", integer)
  }
}
