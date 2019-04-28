package main

import (
	myfmt "fmt" // named imports
)

func main() {
	myfmt.Println("hello")
	myfmt.Printf("The quick brown fox jumped over lazy dogs", 3.14)
	/* go vet main.go
	prints below comments
	# command-line-arguments
	./main.go:9:15: Printf call has arguments but no formatting directives */

}
