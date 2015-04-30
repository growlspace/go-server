package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
	// "github.com/martini-contrib/gzip"
)

// User a struct representing a user
type User struct {
	id       int       `json:"id" xml:"id,attr"`
	username xml.Name  `json:"username" xml:"username"`
	name     string    `json:"name" xml:"name"`
	bio      string    `json:"bio" xml:"bio"`
	created  time.Time `json:"created" xml:"created"`
	updated  time.Time `json:"updated" xml:"updated"`
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
	// add martini-contrib secure?
	// m.Use(gzip.All())

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

	m.Get("/v1/feed/:last_id", func(params martini.Params, rw http.ResponseWriter, r *http.Request, db *sql.DB) {
		query := "SELECT caption FROM posts WHERE post_id - " + params["last_id"] + " > ORDER BY post_id DESC LIMIT 10"
		rows, err := db.Query(query)
		panicIf(err)
		defer rows.Close()
		var id int
		var caption string
		for rows.Next() {
			err := rows.Scan(&id, &caption)
			panicIf(err)

			fmt.Fprintf(rw, "ID: %d\nCaption: %s\n\n", id, caption)
		}

	})

	m.Get("/v1/feed/", func(params martini.Params, rw http.ResponseWriter, r *http.Request, db *sql.DB) {
		query := "SELECT caption FROM posts ORDER BY post_id DESC LIMIT 10"
		rows, err := db.Query(query)
		panicIf(err)
		defer rows.Close()
		var id int
		var caption string
		for rows.Next() {
			err := rows.Scan(&id, &caption)
			panicIf(err)

			fmt.Fprintf(rw, "ID: %d\nCaption: %s\n\n", id, caption)
		}

	})

	m.Get("/v1/feed/:last_id", func(params martini.Params, rw http.ResponseWriter, r *http.Request, db *sql.DB) {
		query := "SELECT caption FROM posts WHERE " + params["last_id"] + " - post_id > 0 ORDER BY post_id DESC LIMIT 10"
		rows, err := db.Query(query)
		panicIf(err)
		defer rows.Close()
		var id int
		var caption string
		for rows.Next() {
			err := rows.Scan(&id, &caption)
			panicIf(err)

			fmt.Fprintf(rw, "ID: %d\nCaption: %s\n\n", id, caption)
		}

	})

	m.Run()
}
