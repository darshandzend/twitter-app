package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func check(err error, line int) {
	if err != nil {
		log.Print(err, line)
		panic(err)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {

	b := make([]byte, r.ContentLength)
	_, err := r.Body.Read(b)
	//	check(err, 21) : handle EOF
	log.Print(string(b))

	var d draft
	err = json.Unmarshal(b, &d)
	check(err, 26)

	err = validateDraft(d)
	check(err, 29)

	err = createDraftDAO(d)
	check(err, 32)

	fmt.Fprintf(w, "saved")

}
