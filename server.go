package main

import (
	"database/sql"

	"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
	// "github.com/martini-contrib/gzip"
)

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

	m.Get("/v1/user/:username", GetUser)
	m.Get("/v1/feed", GetFeed)

	m.Run()
}
