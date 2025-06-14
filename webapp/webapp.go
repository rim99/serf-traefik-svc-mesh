package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getAppName() string {
	name := os.Getenv("APP_NAME")
	if name == "" {
		name = "unset"
	}
    return name
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	helloMsg := "Hello, from " + getAppName() + "!"

	w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintf(w, helloMsg)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	remote := r.URL.Query().Get("remote")
	targetURL := "http://" + remote + ":28080/hello" // the proxy serves traffic to remote on port 28080
	resp, err := http.Get(targetURL)
	if err != nil {
		http.Error(w, "Failed to request target service: " + targetURL, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(body)
}

func main() {
    http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/ping", pingHandler)

    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}

