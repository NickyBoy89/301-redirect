package main

import (
  "net/http"
  "fmt"
  "os"
)

var documentRoot string = "0.0.0.0" // Change this to 127.0.0.1 for localhost

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
  // Redirect the incoming HTTP request
  http.Redirect(w, r, os.Getenv("URL")+r.RequestURI, http.StatusMovedPermanently)
}
func main() {
  if _, present := os.LookupEnv("URL"); present == false {
    panic("URL environment variable must be set for redirect to function")
  }
  // Status message on startup
  fmt.Printf("Server listening on %v:8080\n", documentRoot)

  // Start the HTTP server and redirect all incoming connections
  if err := http.ListenAndServe(documentRoot + ":8080", http.HandlerFunc(redirectToHttps)); err != nil {
    fmt.Printf("Fatal error: %v\n", err)
  }
}
