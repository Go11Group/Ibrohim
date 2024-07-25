CREATE TABLE IF NOT EXISTS user_locations (
    location_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    address VARCHAR(255),
    city VARCHAR(100),
    state VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO user_locations (location_id, user_id, address, city, state, country, postal_code) VALUES
    ('c0a80132-0000-1000-8000-00805f9b34fb', 'c0a80122-0000-1000-8000-00805f9b34fb', '123 Main St', 'Springfield', 'IL', 'USA', '62701'),
    ('c0a80133-0000-1000-8000-00805f9b34fb', 'c0a80123-0000-1000-8000-00805f9b34fb', '456 Elm St', 'Metropolis', 'NY', 'USA', '10001'),
    ('c0a80134-0000-1000-8000-00805f9b34fb', 'c0a80124-0000-1000-8000-00805f9b34fb', '789 Oak St', 'Gotham', 'NJ', 'USA', '07001'),
    ('c0a80135-0000-1000-8000-00805f9b34fb', 'c0a80125-0000-1000-8000-00805f9b34fb', '101 Pine St', 'Star City', 'CA', 'USA', '90001'),
    ('c0a80136-0000-1000-8000-00805f9b34fb', 'c0a80126-0000-1000-8000-00805f9b34fb', '202 Maple St', 'Central City', 'CO', 'USA', '80427'),
    ('c0a80137-0000-1000-8000-00805f9b34fb', 'c0a80127-0000-1000-8000-00805f9b34fb', '303 Cedar St', 'Coast City', 'OR', 'USA', '97034'),
    ('c0a80138-0000-1000-8000-00805f9b34fb', 'c0a80128-0000-1000-8000-00805f9b34fb', '404 Birch St', 'Keystone City', 'KS', 'USA', '66101'),
    ('c0a80139-0000-1000-8000-00805f9b34fb', 'c0a80129-0000-1000-8000-00805f9b34fb', '505 Walnut St', 'National City', 'CA', 'USA', '91950'),
    ('c0a80140-0000-1000-8000-00805f9b34fb', 'c0a80130-0000-1000-8000-00805f9b34fb', '606 Chestnut St', 'Smallville', 'KS', 'USA', '66002'),
    ('c0a80141-0000-1000-8000-00805f9b34fb', 'c0a80131-0000-1000-8000-00805f9b34fb', '707 Spruce St', 'Bludhaven', 'NJ', 'USA', '07002');