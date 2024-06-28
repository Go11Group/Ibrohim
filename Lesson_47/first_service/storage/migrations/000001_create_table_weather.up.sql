CREATE TYPE weather AS enum('sunny', 'rainy', 'snowy', 'stormy', 'windy', 'warm', 'cool', 'cold');
CREATE TABLE IF NOT EXISTS weather_conditions (
    city VARCHAR NOT NULL,
    date DATE DEFAULT CURRENT_DATE,
    weather_type weather,
    temperature INT,
    humidity INT,
    wind_speed INT
);