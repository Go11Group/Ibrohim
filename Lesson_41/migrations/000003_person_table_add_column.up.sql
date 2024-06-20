CREATE TYPE marital_status AS enum ('married', 'not married');
ALTER TABLE person ADD COLUMN marital_status marital_status;