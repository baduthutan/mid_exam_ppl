package main

import (
	"fmt"
	"net/http"
	"datamanager"
)

// Take in an http.ResponseWriter and a pointer to an http.Request struct as arguments. 
func HandleUrl(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		datamanager.DisplayEveryData(w, r)
	case http.MethodPost:
		datamanager.AppendData(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// register the HandleUrl function to the root URL of the server and start the server.
func main() {
	http.HandleFunc("/", HandleUrl)
	fmt.Println("Server started on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}