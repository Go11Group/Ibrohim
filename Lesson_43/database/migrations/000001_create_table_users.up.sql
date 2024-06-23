CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    age INT CHECK (age > 0),
    phone_number VARCHAR(20) UNIQUE
);