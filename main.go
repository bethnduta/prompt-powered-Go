package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

type Response struct {
    Message string `json:"message"`
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(Response{Message: "Hello, World!"})
    })

    http.HandleFunc("/welcome/", func(w http.ResponseWriter, r *http.Request) {
        name := strings.TrimPrefix(r.URL.Path, "/welcome/")
        if name == "" {
            name = "Guest"
        }
        json.NewEncoder(w).Encode(Response{Message: fmt.Sprintf("Welcome, %s!", name)})
    })

    fmt.Println("🌍 Server running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
