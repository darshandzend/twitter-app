package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func createDraftDAO(d Draft) (string, error) {
	con, err := sql.Open("mysql", squser+":"+pwd+"@/"+database)
	check(err)
	defer con.Close()

	q := "insert into draft (id, user_id, text) values (?, ?, ?)"

	res, err := con.Exec(q, d.Id, 0, d.Text)
	check(err)

	lastId, err := res.LastInsertId()
	check(err)

	id := strconv.FormatInt(lastId, 10)

	return id, nil
}

func readAllDAO(user_id int) ([]*Draft, error) {
	con, err := sql.Open("mysql", squser+":"+pwd+"@/"+database)
	check(err)
	defer con.Close()

	q := "select id, text " +
		"from draft " +
		"where user_id=?"

	rows, err := con.Query(q, string(user_id))
	check(err)
	drafts := make([]*Draft, 0, 10)
	var id, text string
	for rows.Next() {
		err = rows.Scan(&id, &text)
		check(err)
		drafts = append(drafts, &Draft{id, string(user_id), text})
	}

	return drafts, nil
}

func readDAO(id string, user_id int) (*Draft, error) {
	con, err := sql.Open("mysql", squser+":"+pwd+"@/"+database)
	if err != nil {
		return nil, err
	}

	defer con.Close()

	q := "select text " +
		"from draft " +
		"where user_id=? " +
		"and id=?"

	row := con.QueryRow(q, string(user_id), id)
	var text string
	err = row.Scan(&text)
	if err != nil {
		return nil, fmt.Errorf("No record with id %s found", id)
	}
	return &Draft{id, string(user_id), text}, nil

}
