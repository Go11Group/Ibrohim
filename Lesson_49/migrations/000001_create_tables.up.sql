CREATE TYPE episode_status AS ENUM ('draft', 'in_progress', 'completed', 'published');
CREATE TABLE podcasts (
    id uuid PRIMARY KEY default gen_random_uuid() not null ,
    user_id uuid not null,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    status episode_status DEFAULT 'draft',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP not null,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP    
);

CREATE TABLE episodes (
      id uuid PRIMARY KEY default gen_random_uuid() not null,
      podcast_id uuid not null ,
      user_id uuid not null ,
      title VARCHAR(100) NOT NULL,
      file_audio bytea NOT NULL,
      description TEXT,
      duration float,
      created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP not null,
      updated_at TIMESTAMP,
      deleted_at TIMESTAMP
);
