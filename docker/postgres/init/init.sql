DROP TABLE IF EXISTS users CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE DATABASE app;

GRANT ALL PRIVILEGES ON DATABASE app TO postgres;

CREATE TABLE  users (
		id                   UUID              PRIMARY KEY                     DEFAULT uuid_generate_v4(),
		user_name            VARCHAR(255) NOT NULL,
		email                VARCHAR(255) UNIQUE,
		hashed_password      bytea NOT NULL,
		salt                 bytea NOT NULL,
		created_at           TIMESTAMP WITH TIME ZONE        NOT NULL DEFAULT NOW(),
		updated_at           TIMESTAMP WITH TIME ZONE        DEFAULT CURRENT_TIMESTAMP
    );
