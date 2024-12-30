package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go API!")
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the API endpoint!")
	})

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
