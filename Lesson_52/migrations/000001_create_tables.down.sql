drop INDEX if exists idx_refresh_tokens_token;
drop INDEX if exists idx_refresh_tokens_user_id;
drop table if exists refresh_tokens;
drop table if exists user_profiles;
drop type if exists user_role;
drop table if exists users;