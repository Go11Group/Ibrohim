CREATE TABLE IF NOT EXISTS schedule (
    bus_id uuid REFERENCES buses(id),
    monday VARCHAR NOT NULL,
    tuesday VARCHAR NOT NULL,
    wednesday VARCHAR NOT NULL,
    thursday VARCHAR NOT NULL,
    friday VARCHAR NOT NULL,
    saturday VARCHAR NOT NULL,
    sunday VARCHAR NOT NULL
);