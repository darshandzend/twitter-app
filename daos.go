package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func createDraftDAO(d Draft) error {
	con, err := sql.Open("mysql", squser+":"+pwd+"@/"+database)
	check(err)
	defer con.Close()

	_, err = con.Exec("insert into draft (draft_id, user_id, text) values (?, ?, ?)", d.Id, 0, d.Text)
	check(err)

	return nil
}

func readAllDAO(user_id int) ([]*Draft, error) {
	con, err := sql.Open("mysql", squser+":"+pwd+"@/"+database)
	check(err)
	defer con.Close()

	rows, err := con.Query(
		"select draft_id, text "+
			"from draft "+
			"where user_id=?",
		string(user_id))
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
