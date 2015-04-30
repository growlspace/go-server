package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/martini"
)

// GetUser gets a user from the url path param based username
func GetUser(params martini.Params, r *http.Request, db *sql.DB) string {

	query := "SELECT user_id,username,real_name,bio FROM users WHERE username LIKE '" + params["username"] + "'"
	rows, err := db.Query(query)
	panicIf(err)
	defer rows.Close()
	var resultUser User
	for rows.Next() {
		err := rows.Scan(&resultUser.id, &resultUser.username, &resultUser.name, &resultUser.bio)
		panicIf(err)
	}
	response, err := json.Marshal(resultUser)
	if err != nil {
		log.Fatal(err)
	}
	return string(response)
}

// GetFeed gets a slice of feed items (up to 10) optionally older than an ID supplied as a query param
func GetFeed(params martini.Params, r *http.Request, db *sql.DB) string {
	lastID := r.URL.Query().Get("last_id")
	sqlQuery := "SELECT caption FROM posts ORDER BY post_id DESC LIMIT 10"
	if lastID != "" {
		sqlQuery = "SELECT caption FROM posts WHERE post_id - " + lastID + " > ORDER BY post_id DESC LIMIT 10"
	}
	rows, err := db.Query(sqlQuery)
	panicIf(err)
	defer rows.Close()

	posts := make([]Post, 10)
	i := 0
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.id, &p.caption)
		panicIf(err)

		posts[i] = p
		i++
	}
	response, err := json.Marshal(posts)
	if err != nil {
		log.Fatal(err)
	}
	return string(response)
}
