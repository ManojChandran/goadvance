package main

import (
  "fmt"
)

func main(){
  items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
  fmt.Println(linearsearch(items, 51))
}

func linearsearch(data []int, key int) bool  {
  for _, item := range data{
    if item == key{
      return true
    }
  }
  return false
}
