package db

import (
	"database/sql"
	"fmt"
)

var Pdb *sql.DB

func init() {
	var err error
	Pdb, err = sql.Open("postgres", "postgres://postgres:duong@localhost/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = Pdb.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
