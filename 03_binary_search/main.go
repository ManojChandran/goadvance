package main

import (
  "fmt"
)

func main(){
  items := []int{ 1, 2, 9, 20, 31, 45, 63, 70, 100 }
  start := 0
  end := len(items) - 1
  fmt.Println(linearsearch(items, 0, start, end))
//  fmt.Println(linearsearch(items, 20, start, end))
//  fmt.Println(linearsearch(items, 63, start, end))
//  fmt.Println(linearsearch(items, 70, start, end))
}

func linearsearch(data []int, key, start, end int) bool  {
  middle := (start + end) / 2
  fmt.Println(data, start, end, data[middle], key)

  if key == data[middle] {
    return true
  }

  if key > data[middle] {
    start = middle + 1
    fmt.Println("1", start, end, key)
  }

  if key < data[middle] {
    end = middle - 1
    fmt.Println("2 ",start, end)
  }

  if start > end {
    return false
  }

  return linearsearch(data, key, start, end)
}
