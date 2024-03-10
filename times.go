package main

import (
	"embed"
	"fmt"
	"net"
	"net/http"
	"time"
)

const (
	port = ":8080"
	host = "localhost"
)

//go:embed index.html styles.css script.js
var staticFiles embed.FS

func main() {
	http.HandleFunc("/time", timeHandler)

	// Serve static files
	http.Handle("/", http.FileServer(http.FS(staticFiles)))

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
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(currentTime))

	// UDP code remains unchanged
	message := []byte(currentTime)
	addr, err := net.ResolveUDPAddr("udp", host+":123")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Failed to establish UDP connection:", err)
		return
	}
	defer conn.Close()
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Failed to send UDP packet:", err)
		return
	}
	fmt.Println("UDP packet sent successfully")
	fmt.Println("Time is:", currentTime)
}
