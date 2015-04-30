package main

import (
	"github.com/codegangsta/martini"
	"net/http"
	"fmt"
	"time"
	"database/sql"
	_ "github.com/lib/pq"
)
func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=passwordlol dbname=postgres sslmode=disable")
	PanicIf(err)
	return db
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	m := martini.Classic()

	m.Map(SetupDB())

	m.Get("/", func(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
		rows, err := db.Query("SELECT * FROM users")
		PanicIf(err)
		defer rows.Close()
		var id, username, password, name, bio string
		var created, updated time.Time
		for rows.Next() {
			err := rows.Scan(&id, &username, &password, &name, &bio, &created, &updated)
			PanicIf(err)

			fmt.Fprintf(rw, "ID: %s\nUser: %s\nPassword: %s\nName: %s\nBio: %s\nCreated: %s\nUpdated: %s\n", id, username, password, name, bio, created, updated)
		}
	})

	m.Run()
}
