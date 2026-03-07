package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DevOps Pipeline Demo App Running 🚀")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on ports 8080")
	http.ListenAndServe(":8080", nil)
}