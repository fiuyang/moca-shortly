CREATE TABLE IF NOT EXISTS urls (
    id SERIAL PRIMARY KEY,
    short_code VARCHAR(125) NOT NULL,
    original_url TEXT NOT NULL,
    clicks INT DEFAULT 0,
    expired_at TIMESTAMP NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NULL
);

ALTER TABLE urls
ADD CONSTRAINT unique_short_code UNIQUE (short_code);