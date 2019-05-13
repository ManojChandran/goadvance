/*
  go methods
*/
package main

import (
  "fmt"
)

type User struct {
  FirstName, LastName string
}

func (u User) Greetings() string  {
  return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main()  {
  u := User {"Manoj", "Chandran"}
  fmt.Println(u.Greetings())
}
