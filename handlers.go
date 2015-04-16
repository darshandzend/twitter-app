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
		log.Print(err)
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

	//TODO: retrieve user, currently default

	var d Draft
	err = json.Unmarshal(b, &d)
	check(err)

	err = validateDraft(d)
	check(err)

	err = createDraftDAO(d)
	check(err)

	fmt.Fprintf(w, "saved")

}

func readAllHandler(w http.ResponseWriter, r *http.Request) {

	/*b, err := readBody(r)
	check(err)*/

	//TODO: retrieve user, currently default
	u := 0

	drafts := make([]*Draft, 0, 10)
	drafts, err := readAllDAO(u) //TODO: Pagination
	check(err)

	resp, err := json.Marshal(drafts)
	check(err)

	fmt.Fprintf(w, string(resp))

}

/*






















*/
