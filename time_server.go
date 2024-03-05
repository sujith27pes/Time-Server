package main

import (
    "fmt"
    "net"
    "net/http"
    "time"
)

const (
    port = ":8080"
    host = "localhost"
)

func main() {
    http.HandleFunc("/time", timeHandler)

    fmt.Printf("Server listening on %s\n", port)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    currentTime := time.Now().Format(time.RFC1123)
    message := []byte(currentTime)

    addr, err := net.ResolveUDPAddr("udp", host+":123")
    if err != nil {
        http.Error(w, "Failed to resolve UDP address", http.StatusInternalServerError)
        return
    }

    conn, err := net.DialUDP("udp", nil, addr)
    if err != nil {
        http.Error(w, "Failed to establish UDP connection", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    _, err = conn.Write(message)
    if err != nil {
        http.Error(w, "Failed to send UDP packet", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "UDP packet sent successfully")
  fmt.Println("Time is: ", currentTime)
}

