# server
back-end server code

nick's required APIs:
- api.growl.space/1.0/login (POST)
- api.growl.space/1.0/feed (GET)
- api.growl.space/1.0/item/{id} (GET, POST)
- api.growl.space/1.0/user/{username} (GET, POST)

expectations:
* POST login with username with credentials (in body), receive auth token
* GET feed, receive list of item [requires auth]
* GET item with id, receive content [requires auth]
* POST item without id, create content [requires auth]
* GET user with username, receive user profile content [requires auth]
* POST user, create user

all in JSON

thx :)
