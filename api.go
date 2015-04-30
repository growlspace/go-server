package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/martini"
)

// GetUser gets a user from the url path param based ID
func GetUser(params martini.Params, r *http.Request, db *sql.DB) string {
	query := "SELECT user_id,username,real_name,bio,created_at,updated_at FROM users WHERE user_id=" + params["id"]
	rows, err := db.Query(query)
	panicIf(err)
	defer rows.Close()
	var resultUser User
	var created, updated time.Time
	for rows.Next() {
		err := rows.Scan(&resultUser.ID, &resultUser.Username, &resultUser.Name, &resultUser.Bio, &created, &updated)
		panicIf(err)
		resultUser.Created = created.Unix()
		resultUser.Updated = updated.Unix()
	}
	fmt.Println(resultUser)
	response, err := json.Marshal(resultUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response))
	return string(response)
}

// GetFeed gets a slice of feed items (up to 10) optionally older than an ID supplied as a query param
func GetFeed(params martini.Params, r *http.Request, db *sql.DB) string {
	lastID := r.URL.Query().Get("last_id")
	sqlQuery := "SELECT * FROM posts ORDER BY post_id DESC LIMIT 10"
	if lastID != "" {
		sqlQuery = "SELECT * FROM posts WHERE post_id - " + lastID + " > ORDER BY post_id DESC LIMIT 10"
	}
	rows, err := db.Query(sqlQuery)
	panicIf(err)
	defer rows.Close()

	posts := make([]Post, 10)
	i := 0
	var created, updated time.Time
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.CreatedBy, &p.Caption, &p.Audio, &created, &updated)
		panicIf(err)
		p.Created = created.Unix()
		p.Updated = updated.Unix()
		posts[i] = p
		i++
	}

	response, err := json.Marshal(posts)

	if err != nil {
		log.Fatal(err)
	}
	return string(response)
}

// GetItem gets an item by ID
func GetItem(params martini.Params, r *http.Request, db *sql.DB) string {
	query := "SELECT * FROM posts WHERE post_id=" + params["id"]
	rows, err := db.Query(query)
	panicIf(err)
	defer rows.Close()
	var p Post
	var created, updated time.Time
	for rows.Next() {
		err := rows.Scan(&p.ID, &p.CreatedBy, &p.Caption, &p.Audio, &created, &updated)
		panicIf(err)
		p.Created = created.Unix()
		p.Updated = updated.Unix()
	}
	fmt.Println(p)
	response, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response))
	return string(response)
}
