CREATE SCHEMA posts AUTHORIZATION glog;

CREATE TABLE post (
    post_id SERIAL PRIMARY KEY,
    post_date TIMESTAMP,
    post_author int,
    title varchar(512),
    body_text TEXT
);

CREATE TABLE authors (
    author_id SERIAL PRIMARY KEY,
    author_user varchar(32),
    author_name varchar(512)
);

CREATE TABLE users (
    username varchar(32) PRIMARY KEY,
    password bytea
)