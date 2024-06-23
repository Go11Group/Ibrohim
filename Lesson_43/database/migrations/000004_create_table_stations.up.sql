CREATE TABLE IF NOT EXISTS stations (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL
);