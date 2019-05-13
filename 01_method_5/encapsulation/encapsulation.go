package encapsulation

import (
    "fmt"
)

// Encapsulation struct can be exported outside of this package
type Encapsulation struct{}

// Expose method can be exported outside of this package
func (e Encapsulation) Expose ()  {
  fmt.Println("AHHHH! I am exposed")
}

// hide method can only be used within this package
func (e Encapsulation) hide()  {
  fmt.Println("SHHHHH... this is super secret")
}

// unhide uses the unexpected hide function
func (e Encapsulation) Unhide() {
  e.hide()
  fmt.Println(".....jk")
}
