CREATE TYPE transaction_type AS ENUM ('credit', 'debit');
CREATE TABLE IF NOT EXISTS transactions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    card_id uuid REFERENCES cards(id),
    amount INT CHECK (amount > 0),
    terminal_id uuid REFERENCES terminals(id) DEFAULT NULL,
    type transaction_type
);