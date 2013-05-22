package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world!")
}

func main() {
    fmt.Printf("Hello world!\n")
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler)
}