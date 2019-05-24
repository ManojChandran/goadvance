// Go channels
/*
Channels are on of the synchronization
primitives drived in go from Hoare's
CSP
*/
package main

import (
    "fmt"
)

func main() {
  // Channel declaration
  var receiveChan <-chan interface{} // unidirectional receiver
  var sendChan chan<- interface{} // unidirectional Sender
  dataStream := make(chan interface{}) // Bidirectional
  // valid statement
  receiveChan = dataStream
  sendChan = dataStream
  fmt.Println(receiveChan)
  fmt.Println(sendChan)
  fmt.Println(dataStream)  
}
