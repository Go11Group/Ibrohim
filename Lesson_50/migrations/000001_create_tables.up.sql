CREATE TYPE collaboration_role AS ENUM ('owner', 'collaborator');
CREATE TYPE invitation_status AS ENUM ('pending', 'accepted', 'declined');

CREATE TABLE collaborations (
    id uuid PRIMARY KEY gen_random_uuid(),
    podcast_id uuid,
    user_id uuid,
    role collaboration_role DEFAULT 'collaborator',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE invitations (
    id uuid PRIMARY KEY gen_random_uuid(),
    podcast_id uuid,
    inviter_id uuid,
    invitee_id uuid,
    status invitation_status DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE comments (
    id uuid PRIMARY KEY gen_random_uuid(),
    podcast_id uuid,
    episode_id uuid,
    user_id uuid not null,
    content TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
