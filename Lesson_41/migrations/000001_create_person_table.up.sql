CREATE TABLE IF NOT EXISTS person (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    age INTEGER,
    is_married BOOLEAN
);