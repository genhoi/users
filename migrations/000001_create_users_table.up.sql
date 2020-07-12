CREATE TABLE IF NOT EXISTS users(
   id BIGINT PRIMARY KEY,
   login VARCHAR (50) UNIQUE NOT NULL,
   full_name VARCHAR (120) NOT NULL
);

CREATE INDEX users_full_name_fts_idx ON users USING GIN (to_tsvector('russian', full_name));
