-- goose for db migrations
-- goose up
-- up is for adding a new migration

CREATE TABLE users (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email      VARCHAR(255) UNIQUE NOT NULL,
    password   VARCHAR(255) NOT NULL,
    role       TEXT NOT NULL DEFAULT 'user',
    verified   BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

-- goose down
-- down is for removing a migration

DROP TABLE IF EXISTS users;