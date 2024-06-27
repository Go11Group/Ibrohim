CREATE TABLE IF NOT EXISTS buses (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    number VARCHAR UNIQUE NOT NULL,
    capacity INT CHECK (capacity > 0)
);