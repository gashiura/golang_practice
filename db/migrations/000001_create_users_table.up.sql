BEGIN;
CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   uid VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,
   created_at DATETIME NOT NULL,
   updated_at DATETIME NOT NULL
);
CREATE INDEX users_uid_index ON users(uid);
COMMIT;
