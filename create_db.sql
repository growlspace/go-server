DROP SCHEMA public CASCADE;

CREATE SCHEMA public;

CREATE SEQUENCE user_id_seq;

CREATE TABLE users (user_id int PRIMARY KEY DEFAULT nextval('user_id_seq'), username varchar(10), password_hash text, password_salt text, real_name text, bio text, created_at timestamp, updated_at timestamp);

CREATE SEQUENCE audio_id_seq;

CREATE TABLE audio (audio_id int PRIMARY KEY DEFAULT nextval('audio_id_seq'), audio_filename text NOT NULL, created_at timestamp, updated_at timestamp);

CREATE SEQUENCE post_id_seq;

CREATE TABLE posts (post_id int PRIMARY KEY DEFAULT nextval('post_id_seq'), user_id int references users(user_id), caption text, audio_id int references audio(audio_id), created_at timestamp, updated_at timestamp);

INSERT INTO "users" (username, password_hash, password_salt, real_name, bio, created_at, updated_at) VALUES ('test', 'passwordlol', 'lol', 'John Smith', 'hes a mean one', current_timestamp, current_timestamp);

INSERT INTO "users" (username, password_hash, password_salt, real_name, bio, created_at, updated_at) VALUES ('test2', 'passwordlol', 'lol', 'Jane Smith', 'shes a mean one', current_timestamp, current_timestamp);

INSERT INTO "audio" (audio_filename, created_at, updated_at) VALUES ('this_is_a_file.3gp', current_timestamp, current_timestamp);

INSERT INTO "audio" (audio_filename, created_at, updated_at) VALUES ('this_is_another_file.3gp', current_timestamp, current_timestamp);

INSERT INTO "posts" (user_id, caption, audio_id, created_at, updated_at) VALUES (1, 'ay girl', 1, current_timestamp, current_timestamp);

INSERT INTO "posts" (user_id, caption, audio_id, created_at, updated_at) VALUES (2, 'ay girl', 2, current_timestamp, current_timestamp);
