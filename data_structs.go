package main

import (
	"time"
)

// User a struct representing a user
type User struct {
	id       int
	username string
	name     string
	bio      string
	created  time.Time
	updated  time.Time
}

// Post is a struct representing a post by a user
type Post struct {
	id       int
	caption  string
	user     User
	audioURL string
	created  time.Time
	updated  time.Time
}
