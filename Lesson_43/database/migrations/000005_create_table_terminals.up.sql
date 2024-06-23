CREATE TABLE IF NOT EXISTS terminals (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    station_id uuid REFERENCES stations(id)
);