package main

import (
  "fmt"
  "github.com/braintree/manners"
  "github.com/gorilla/mux"
  "github.com/jcarley/basic_s3_uploader/backends/golang/handlers"
  "log"
  "net/http"
  "os"
)

const port = ":8090"

var r *mux.Router

// This filter enables messing with the request/response before and after the normal handler
func filter(w http.ResponseWriter, req *http.Request) {
  r.ServeHTTP(w, req) // calls the normal handler
  log.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.URL)
}

func main() {
  wd, _ := os.Getwd()
  println("Working directory", wd)

  r = mux.NewRouter()
  handlers.Install(r)
  http.HandleFunc("/", filter)

  fmt.Println("Running on " + port)

  server := manners.NewServer()
  server.ListenAndServe(port, nil)
}
