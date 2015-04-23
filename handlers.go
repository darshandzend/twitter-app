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

func makeHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Recieved %s %s", r.Method, r.URL)
		err := f(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Handling %q: %v", r.RequestURI, err)
		}
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

func createHandler(w http.ResponseWriter, r *http.Request) error {

	b, err := readBody(r)
	check(err)

	//TODO: retrieve user, currently default

	var d Draft
	err = json.Unmarshal(b, &d)
	check(err)

	err = validateDraft(d)
	check(err)

	id, err := createDraftDAO(d)
	check(err)

	resp, err := json.Marshal(id)

	fmt.Fprintf(w, string(resp))

	return err

}

func readAllHandler(w http.ResponseWriter, r *http.Request) error {

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

	return err
}

func readHandler(w http.ResponseWriter, r *http.Request) error {

	u := 0

	//validate request
	id := r.URL.Query()["id"][0]

	//validate

	d, err := readDAO(id, u)
	if err != nil {
		return err
	}
	resp, err := json.Marshal(d)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, string(resp))
	return err
}

/*






















*/
