// go reflections
// Type Assertion
package main

import (
	"fmt"
)

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	//  func Println(a ...interface{}) (n int, err error)
	Println(true)
	Println(2010)
	Println(9.15)
	Println(7 + 7i)
	Println("Hello world")
	Println(ID("19520925"))
	Println([5]byte{})
	Println([]byte{})
	Println(map[string]int{})
	Println(Person{name: "Jane Doe"})
	Println(make(chan int))

}

func Println(x interface{}) {
	if v, ok := x.(ID); ok{
	fmt.Printf("x has type ID, which i defined. The value is :%v\n",v)
	}
	fmt.Printf("type is %T\n", x)
}
