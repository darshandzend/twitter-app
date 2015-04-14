package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func createDraftDAO(d draft) error {
	con, err := sql.Open("mysql", squser+":"+pwd+"@/"+database)
	check(err, 13)
	defer con.Close()

	_, err = con.Exec("insert into draft (draft_id, user_id, text) values (?, ?, ?)", d.Id, 0, d.Text)
	check(err, 17)

	return nil
}
