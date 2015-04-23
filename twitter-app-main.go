package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/create", makeHandler(createHandler))
	http.HandleFunc("/readall", makeHandler(readAllHandler))
	http.HandleFunc("/read", makeHandler(readHandler))

	http.ListenAndServe(":8080", nil)
}
