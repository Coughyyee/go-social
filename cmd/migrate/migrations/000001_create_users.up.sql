CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    email citext UNIQUE NOT NULL, -- case insensitive extension: ex@mail.com == EX@mail.com
    username varchar(255) UNIQUE NOT NULL,
    password bytea NOT NULL,
    created_at timestamp(0) with time zone NOT NULL default NOW()
);
