CREATE TABLE IF NOT EXISTS users (
     id SERIAL PRIMARY KEY,
     name TEXT NOT NULL,
     email TEXT NOT NULL,
     created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
     updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
     deleted_at TIMESTAMPTZ NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email_active
    ON users(email)
    WHERE deleted_at IS NULL;

INSERT INTO users (name, email)
VALUES
    ('Alnur', 'alnur@example.com'),
    ('Test User', 'test@example.com');