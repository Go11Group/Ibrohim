CREATE TABLE IF NOT EXISTS cards (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    number VARCHAR UNIQUE NOT NULL,
    user_id uuid REFERENCES users(id)
);