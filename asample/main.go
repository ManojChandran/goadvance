package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')

    fmt.Println(text)
    fmt.Print("Enter text: ")
    text2, _ := reader.ReadString('\n')
    fmt.Println(text2)
}
