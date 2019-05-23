/*
  go methods
*/
package main

import (
  "fmt"
)
// one way to declare a user-defined type in Go.
type User struct {
  FirstName, LastName string
}
// Methods provide a way to add behavior to user-defined types.
func (u User) Greetings() string  {
  return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main()  {
  u := User {"Manoj", "Chandran"}
  fmt.Println(u.Greetings())
}
