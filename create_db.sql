CREATE SEQUENCE user_id_seq;

CREATE TABLE users (user_id int PRIMARY KEY DEFAULT nextval('user_id_seq'), username varchar(10), password_hash text, password_salt text, real_name text, bio text, created_at timestamp, updated_at timestamp);

CREATE SEQUENCE audio_id_seq;

CREATE TABLE audio (audio_id int primary key, user_id int references users(user_id), audio_url text NOT NULL, created_at timestamp, updated_at timestamp);

CREATE SEQUENCE post_id_seq;

CREATE TABLE posts (post_id int PRIMARY KEY, user_id int references users(user_id), caption text, audio_id int references audio(audio_id), created_at timestamp, updated_at timestamp);

INSERT INTO "users" (username, password, real_name, bio, created_at, updated_at) VALUES ('test', 'passwordlol', 'John Smith', 'hes a mean one', current_timestamp, current_timestamp);

INSERT INTO "users" (username, password, real_name, bio, created_at, updated_at) VALUES ('test2', 'passwordlol', 'Jane Smith', 'shes a mean one', current_timestamp, current_timestamp);
