package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Jenkinsmaster Backend World!")
}

func main() {
    http.HandleFunc("/", helloWorld)
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed: %s", err)
    }
}
