package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

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
	log.Print("Request body:\n %s", string(b))
	return b, nil
}

func createHandler(w http.ResponseWriter, r *http.Request) error {

	b, err := readBody(r)
	if err != nil {
		log.Fatal(err)
		return err
	}

	//TODO: retrieve user, currently default
	u := 0

	var d Draft
	err = json.Unmarshal(b, &d)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = validateDraft(d)
	if err != nil {
		return err
	}

	id, err := createDraftDAO(d, strconv.Itoa(u))
	if err != nil {
		return err
	}

	resp, err := json.Marshal(id)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, string(resp))
	return err

}

func readAllHandler(w http.ResponseWriter, r *http.Request) error {

	//TODO: retrieve user, currently default
	u := 0

	drafts := make([]*Draft, 0, 10)
	drafts, err := readAllDAO(strconv.Itoa(u)) //TODO: Pagination
	if err != nil {
		return err
	}

	resp, err := json.Marshal(drafts)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, string(resp))
	return err
}

func readHandler(w http.ResponseWriter, r *http.Request) error {

	u := 0

	id := r.URL.Query()["id"][0]

	//validate

	d, err := readDAO(id, strconv.Itoa(u))
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
