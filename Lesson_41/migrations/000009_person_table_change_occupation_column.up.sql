ALTER TABLE person DROP COLUMN occupation;
ALTER TABLE person ADD COLUMN occupation_id INT REFERENCES occupation(id);