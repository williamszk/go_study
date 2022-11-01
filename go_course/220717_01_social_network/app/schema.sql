-- CockroachDB schema
-- app name is nakama


DROP DATABASE IF EXISTS nakama CASCADE;

CREATE DATABASE IF NOT EXISTS nakama;

SET DATABASE = nakama;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    username VARCHAR NOT NULL UNIQUE
);


