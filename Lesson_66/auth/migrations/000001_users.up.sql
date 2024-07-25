CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR(255),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    phone VARCHAR(255),
    image VARCHAR(255)[],
    role VARCHAR(50) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

INSERT INTO users (id, full_name, username, email, password_hash, phone, image, role)
VALUES
    ('c0a80122-0000-1000-8000-00805f9b34fb', 'John Doe', 'johndoe', 'johndoe@example.com', '$2b$12$4MLlI.VpV5OeD7v8E9B5ouDl1zGyW1BWmL/6NdKUuEKkzXn3/7N96', '1234567890', ARRAY['default_image.png'], 'user'), -- password: Password123
    ('c0a80123-0000-1000-8000-00805f9b34fb', 'Jane Smith', 'janesmith', 'janesmith@example.com', '$2b$12$t5B9DlFe4Ojaf10y0E/KHe0eYlJTHdYFB0Y9T/suhmK6YwB1LTiZa', '0987654321', ARRAY['default_image.png'], 'user'), -- password: MySecret123
    ('c0a80124-0000-1000-8000-00805f9b34fb', 'Alice Johnson', 'alicej', 'alicej@example.com', '$2b$12$2cPeW3EKK5lHE4Ipt95S5eeJ2A.Rlq6ZCpHhx5JZnB2.PuGH1GmGq', '5551234567', ARRAY['default_image.png'], 'admin'), -- password: AdminPass!
    ('c0a80125-0000-1000-8000-00805f9b34fb', 'Bob Brown', 'bobbrown', 'bobbrown@example.com', '$2b$12$KqJW1jDzxP7/jJrBp6NRYuz6HyXx33F5v2Wpm8NG1n7H1mJhDKD76', '5557654321', ARRAY['default_image.png'], 'user'), -- password: BobSecure1
    ('c0a80126-0000-1000-8000-00805f9b34fb', 'Charlie Green', 'charlieg', 'charlieg@example.com', '$2b$12$vTCHosOVJcMgiqM5eAGKi.p.LdcqG3/LZgxoFEGp7V7lGInN0zQMS', '5551112222', ARRAY['default_image.png'], 'user'), -- password: Green12345
    ('c0a80127-0000-1000-8000-00805f9b34fb', 'David White', 'davidw', 'davidw@example.com', '$2b$12$zclfjx7t2O3oBL7QL7oZFuZp8.bYQsAjsa8uD1Mh3KHVpxZXvgEGK', '5553334444', ARRAY['default_image.png'], 'user'), -- password: WhiteSecure!
    ('c0a80128-0000-1000-8000-00805f9b34fb', 'Eve Black', 'eveb', 'eveb@example.com', '$2b$12$PzkFVlWWK6fKpLv8Y5A.OOX4hsoYc4z4x6DbYYJeTVogNhyG3aTwq', '5555556666', ARRAY['default_image.png'], 'user'), -- password: BlackEve99
    ('c0a80129-0000-1000-8000-00805f9b34fb', 'Frank Blue', 'frankb', 'frankb@example.com', '$2b$12$QK8J7P/Ql39nMtPDiY9xf.VQoL1bNp9fj1sS4.6lh1JhJqZc06dEO', '5557778888', ARRAY['default_image.png'], 'user'), -- password: BlueFrank77
    ('c0a80130-0000-1000-8000-00805f9b34fb', 'Grace Yellow', 'gracey', 'gracey@example.com', '$2b$12$R6J6y4O4ZpnhGbEoUBZCGOhBzTDoPMRlxJ/LlyZmGshN8OHK8Y38G', '5559990000', ARRAY['default_image.png'], 'user'), -- password: YellowGrace12
    ('c0a80131-0000-1000-8000-00805f9b34fb', 'Hank Red', 'hankr', 'hankr@example.com', '$2b$12$7rGXeWb1RA0Z20A7GcF2O.O2aCvo56Y0p9FNOQrm9mF48M/guBOna', '5552223333', ARRAY['default_image.png'], 'user'); -- password: RedHank45
