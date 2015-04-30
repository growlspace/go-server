package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
)

// User a struct representing a user
type User struct {
	id                  int
	username, name, bio string
	created, updated    time.Time
}

func setupDB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=passwordlol dbname=postgres sslmode=disable")
	panicIf(err)
	return db
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	m := martini.Classic()

	m.Map(setupDB())

	m.Get("/v1/user/:username", func(params martini.Params, rw http.ResponseWriter, r *http.Request, db *sql.DB) {
		query := "SELECT user_id,username,real_name,bio FROM users WHERE username LIKE '" + params["username"] + "'"
		rows, err := db.Query(query)
		panicIf(err)
		defer rows.Close()
		var id int
		var username, name, bio string
		for rows.Next() {
			err := rows.Scan(&id, &username, &name, &bio)
			panicIf(err)

			fmt.Fprintf(rw, "ID: %d\nUser: %s\nName: %s\nBio: %s\n\n", id, username, name, bio)
		}

	})

	m.Run()
}
