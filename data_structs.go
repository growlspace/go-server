package main

// User a struct representing a user
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"real_name"`
	Bio      string `json:"bio"`
	Created  int64  `json:"created_at"`
	Updated  int64  `json:"updated_at"`
}

// AudioRecord contain the filename for the audio
type AudioRecord struct {
	ID       int    `json:"id"`
	Filename string `json:"filename"`
	Created  int64  `json:"created_at"`
	Updated  int64  `json:"updated_at"`
}

// Post is a struct representing a post by a user
type Post struct {
	ID        int    `json:"id"`
	Caption   string `json:"caption"`
	CreatedBy int    `json:"created_by"`
	Audio     int    `json:"audio"`
	Created   int64  `json:"created_at"`
	Updated   int64  `json:"updated_at"`
}
