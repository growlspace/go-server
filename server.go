package main

import (
	"database/sql"

	"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
	// "github.com/martini-contrib/gzip"
)

func setupDB() *sql.DB {
	db, err := sql.Open("postgres", "user="+DBUser+" password="+DBPassword+" dbname=postgres sslmode=disable")
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

	// needs auth or something I think
	m.Get("/v1/user/:id", GetUser)
	// m.Post("/v1/user", AddUser)
	m.Get("/v1/feed", GetFeed)
	m.Get("/v1/item/:id", GetItem)
	// m.Post("/v1/item", AddItem)

	m.Run()
}
