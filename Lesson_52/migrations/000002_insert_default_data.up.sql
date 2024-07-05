INSERT INTO users (id, username, email, password_hash, created_at) VALUES
('c4c8d9e1-0f0b-4b90-b5e3-3f3e2b9a81e1', 'johndoe', 'john.doe@email.com', 'hashed_password_1', NOW()),
('2f15c0f5-6c44-4a23-85a6-7fbb217a628f', 'jane_smith', 'jane.smith@email.com', 'hashed_password_2', NOW()),
('5033b4b5-7d24-4751-98ae-05ee54c31d3d', 'mike_johnson', 'mike.johnson@email.com', 'hashed_password_3', NOW()),
('3c1e2b2e-58d0-40e3-af2c-4d017f9e7c18', 'sarah_williams', 'sarah.williams@email.com', 'hashed_password_4', NOW()),
('5424f49b-e40b-4df2-b0ec-1d87cf57b117', 'david_brown', 'david.brown@email.com', 'hashed_password_5', NOW()),
('eaf3f8f7-ee27-45e2-9460-06de9c7d0c3f', 'emily_davis', 'emily.davis@email.com', 'hashed_password_6', NOW()),
('fbafaf8a-6192-4d4f-bc9d-7f0866e81bf6', 'chris_wilson', 'chris.wilson@email.com', 'hashed_password_7', NOW()),
('bfb5e4de-b54d-48e4-899f-0144ea2c7e64', 'lisa_taylor', 'lisa.taylor@email.com', 'hashed_password_8', NOW()),
('e6b46b06-6e27-4c57-a62d-cb80f8b35e94', 'robert_anderson', 'robert.anderson@email.com', 'hashed_password_9', NOW()),
('4e5a4e7c-d73e-4eb4-915e-23f01646c775', 'amanda_thomas', 'amanda.thomas@email.com', 'hashed_password_10', NOW());

INSERT INTO user_profiles (user_id, full_name, bio, role, location, avatar_image, website) VALUES
('c4c8d9e1-0f0b-4b90-b5e3-3f3e2b9a81e1', 'John Doe', 'Passionate guitarist and songwriter', 'musician', 'New York, NY', NULL, 'www.johndoemusic.com'),
('2f15c0f5-6c44-4a23-85a6-7fbb217a628f', 'Jane Smith', 'Avid music enthusiast and concert-goer', 'listener', 'Los Angeles, CA', NULL, NULL),
('5033b4b5-7d24-4751-98ae-05ee54c31d3d', 'Mike Johnson', 'Electronic music producer and DJ', 'producer', 'Miami, FL', NULL, 'www.mikejohnsonmusic.com'),
('3c1e2b2e-58d0-40e3-af2c-4d017f9e7c18', 'Sarah Williams', 'Classical pianist and music teacher', 'musician', 'Boston, MA', NULL, 'www.sarahwilliamspiano.com'),
('5424f49b-e40b-4df2-b0ec-1d87cf57b117', 'David Brown', 'Hip-hop artist and beatmaker', 'musician', 'Atlanta, GA', NULL, 'www.davidbrownmusic.com'),
('eaf3f8f7-ee27-45e2-9460-06de9c7d0c3f', 'Emily Davis', 'Indie folk singer-songwriter', 'musician', 'Portland, OR', NULL, 'www.emilydavismusic.com'),
('fbafaf8a-6192-4d4f-bc9d-7f0866e81bf6', 'Chris Wilson', 'Music blog writer and critic', 'listener', 'Chicago, IL', NULL, 'www.chriswilsonreviews.com'),
('bfb5e4de-b54d-48e4-899f-0144ea2c7e64', 'Lisa Taylor', 'Jazz vocalist and composer', 'musician', 'New Orleans, LA', NULL, 'www.lisataylorjazz.com'),
('e6b46b06-6e27-4c57-a62d-cb80f8b35e94', 'Robert Anderson', 'Record producer and studio owner', 'producer', 'Nashville, TN', NULL, 'www.andersonstudios.com'),
('4e5a4e7c-d73e-4eb4-915e-23f01646c775', 'Amanda Thomas', 'Pop music fan and playlist curator', 'listener', 'Seattle, WA', NULL, NULL);