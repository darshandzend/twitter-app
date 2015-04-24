package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func createDraftDAO(d Draft, u_id string) (string, error) {
	con, err := sql.Open("mysql", squser+":"+pwd+"@/"+database)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer con.Close()

	q := "insert into draft (id, user_id, text) values (?, ?, ?)"
	res, err := con.Exec(q, d.Id, u_id, d.Text)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	id := strconv.FormatInt(lastId, 10)

	return id, nil
}

func readAllDAO(u_id string) ([]*Draft, error) {
	con, err := sql.Open("mysql", squser+":"+pwd+"@/"+database)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.Close()

	q := "select id, text " +
		"from draft " +
		"where user_id=?"

	rows, err := con.Query(q, u_id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	drafts := make([]*Draft, 0, 10)
	var id, text string
	for rows.Next() {
		err = rows.Scan(&id, &text)
		if err != nil {
			log.Fatal(err)
		}
		drafts = append(drafts, &Draft{id, u_id, text})
	}

	if err := rows.Err(); err != nil {
		rows.Close()
		log.Fatal(err)
		return nil, err
	}

	return drafts, nil
}

func readDAO(id string, u_id string) (*Draft, error) {
	con, err := sql.Open("mysql", squser+":"+pwd+"@/"+database)
	if err != nil {
		return nil, err
	}

	defer con.Close()

	q := "select text " +
		"from draft " +
		"where user_id=? " +
		"and id=?"

	var text string
	err = con.QueryRow(q, u_id, id).Scan(&text)
	if err != nil { //&& err != sql.ErrNoRows {
		return nil, fmt.Errorf("No record with id %s found", id)
	}
	return &Draft{id, u_id, text}, nil

}
