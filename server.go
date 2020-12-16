
package main

import (
  "encoding/json"
  "net/http"
  "log"
  "fmt"
  "os"
)

type Profile struct {
  Name    string
  Hobbies []string
}

func main() {
  http.HandleFunc("/", myJson)
  fmt.Printf("Starting server at port 8080\n")
  
  port := os.Getenv("PORT")

  if len(port) == 0 {
	port = "8080"
  }

  if err := http.ListenAndServe(":"+port, nil); err != nil {
      log.Fatal(err)
  }
}

func myJson(w http.ResponseWriter, r *http.Request) {
  profile := Profile{"Alex", []string{"snowboarding", "programming"}}

  js, err := json.Marshal(profile)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
  
}