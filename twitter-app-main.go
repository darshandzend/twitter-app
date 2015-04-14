package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/create", createHandler)

	http.ListenAndServe(":8080", nil)
}
