package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
