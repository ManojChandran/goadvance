package main

import (
  "fmt"
  "os"
  "net/http"
)

func main(){

  http.HandleFunc("/", homepage)
  // Retrieves the port from the environment
  http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func homepage(res http.ResponseWriter, req *http.Request)  {
  if req.URL.Path != "/" {
    http.NotFound(res, req)
    return
  }
  fmt.Fprint(res, "The homepage.")
}
