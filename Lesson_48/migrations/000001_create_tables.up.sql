-- Create users table
CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() not null,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Create user_role type
CREATE TYPE user_role AS ENUM ('musician', 'listener', 'producer');

-- Create user_profiles table
CREATE TABLE user_profiles (
    user_id uuid PRIMARY KEY REFERENCES users(id),
    full_name VARCHAR(100),
    bio TEXT,
    role user_role,
    location VARCHAR(100),
    avatar_image bytea,
    website VARCHAR(255)
);