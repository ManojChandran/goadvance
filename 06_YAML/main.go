package main

import (
  "fmt"
  "github.com/kylelemons/go-gypsy/yaml"
)

func main(){
  config, err := yaml.ReadFile("config.yaml")
  if err != nil {
    fmt.Println("Error :", err)
  }
  fmt.Println(config.Get("path"))
}
