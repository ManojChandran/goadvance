package main

import (
  "fmt"
)

func main(){
  items := []int{95,78,46,58,45,86,99,251,320}
  fmt.Println("unsorted -> ", items)
  sortItems(items)
  fmt.Println("sorted -> ", items)
}

func sortItems(items []int)  {
  n := len(items)
  for i := 0; i < n-i; i++{
    for j := 0; j < n-j; j++ {
      if items[j] > items[j+1]{
        items[j], items[j+1] = items[j+1], items[j]
      }
    }
  }
}
