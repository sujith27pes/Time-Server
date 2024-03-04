package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Local()
    fmt.Fprintf(w, "The current time is: %v", currentTime)
}

func main() {
    // Serve static files
    http.Handle("/", http.FileServer(http.Dir(".")))
    
    // Handle the time endpoint
    http.HandleFunc("/time", handler)
    
    fmt.Println("Starting time server...")
    http.ListenAndServe(":8080", nil)
}
