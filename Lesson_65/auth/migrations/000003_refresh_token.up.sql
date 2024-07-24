CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY default gen_random_uuid(),
    user_id UUID references users(id),
    token text not null,
    expires_at TIMESTAMP not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);