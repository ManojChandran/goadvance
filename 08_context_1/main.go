// go context example
package main

import (
  "fmt"
  "time"
  "context"
  "os"
  "bufio"
)

func main()  {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)

  go func(){
    s := bufio.NewScanner(os.Stdin)
    s.Scan()
    cancel()
  }()

  sleepAndTalk(ctx, 5 * time.Second, "hello")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string)  {
  time.Sleep(d)
  fmt.Println(msg)
}
