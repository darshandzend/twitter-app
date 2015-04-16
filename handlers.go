package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Print(err, line)
		panic(err)
	}
}

func readBody(r *http.Request) ([]byte, error) {
	b := make([]byte, r.ContentLength)
	_, err := r.Body.Read(b)
	if err != nil && err != io.EOF {
		return nil, err
	}
	log.Print(string(b))
	return b, nil
}

func createHandler(w http.ResponseWriter, r *http.Request) {

	b, err := readBody(r)
	check(err)

	var d draft
	err = json.Unmarshal(b, &d)
	check(err)

	err = validateDraft(d)
	check(err)

	err = createDraftDAO(d)
	check(err)

	fmt.Fprintf(w, "saved")

}
