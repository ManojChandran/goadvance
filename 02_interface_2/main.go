// Sample program to show how polymorphic behavior with interfaces.
package main

import (
  "fmt"
)
// notifier is an interface that defines notification
// type behavior
type notifier interface {
  notify()
}

// user defines a user in the program.
type user struct {
  name string
  email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify()  {
  fmt.Printf("Sending user email %s<%s>\n",
  u.name,
  u.email)
}

// admin defines a admin in the program.
type admin struct {
  name string
  email string
}

// notify implements the notifier interface with a pointer receiver.
func (a *admin) notify()  {
  fmt.Printf("Sending admin email %s<%s>\n",
  a.name,
  a.email)
}

func main() {
  // Create a user value and pass it to sendNotification.
  bill := user{"Bill", "bill@email.com"}
  sendNotifier(&bill)

  // Create an admin value and pass it to sendNotification.
  lisa := admin{"Lisa", "lisa@email.com"}
  sendNotifier(&lisa)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotifier(n notifier)  {
  n.notify()
}
