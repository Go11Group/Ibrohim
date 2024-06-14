CREATE DATABASE language_learning_app; -- creating a separate database
CREATE TABLE users (
    user_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    birthday TIMESTAMP NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0); -- creating a table for users
CREATE TABLE courses (
    course_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(100) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0); -- creating a table for courses
CREATE TABLE lessons (
    lesson_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id uuid REFERENCES courses(course_id),
    title VARCHAR(100) NOT NULL UNIQUE,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0); -- creating a table for lessons
CREATE TABLE enrollments (
    enrollment_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid REFERENCES users(user_id),
    course_id uuid REFERENCES courses(course_id),
    enrollment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0,
    CONSTRAINT unique_user_course_combination UNIQUE (user_id, course_id)); -- creating a table for enrollments


INSERT INTO users (name, birthday, email, password)
VALUES
    ('Alice Johnson', '1992-04-10', 'alice.johnson@example.com', 'password123'),
    ('Bob Williams', '1985-09-25', 'bob.williams@example.com', 'securepass'),
    ('Charlie Brown', '1998-12-15', 'charlie.brown@example.com', 'pass1234'),
    ('Diana Martinez', '1987-07-03', 'diana.martinez@example.com', 'letmein'),
    ('Ella Davis', '1994-01-20', 'ella.davis@example.com', 'p@ssw0rd'),
    ('Frank Wilson', '1983-11-12', 'frank.wilson@example.com', 'welcome123'),
    ('Grace Thompson', '1990-06-30', 'grace.thompson@example.com', 'password123'),
    ('Henry Garcia', '1996-08-05', 'henry.garcia@example.com', 'securepass'),
    ('Ivy Rodriguez', '1989-03-18', 'ivy.rodriguez@example.com', 'pass1234'),
    ('Jack Smith', '1993-05-08', 'jack.smith@example.com', 'letmein'); -- inserting 10 users

INSERT INTO courses (title, description)
VALUES
    ('Spanish for Beginners', 'A beginner-level course covering basic Spanish vocabulary and grammar.'),
    ('French Conversation', 'Improve your conversational skills in French with this intermediate-level course.'),
    ('German Grammar Essentials', 'Learn essential German grammar rules and structures in this comprehensive course.'),
    ('Italian Pronunciation Mastery', 'Master Italian pronunciation and phonetics with this specialized course.'),
    ('Mandarin Chinese for Travelers', 'Learn practical Mandarin Chinese phrases and expressions for travelers.'),
    ('Japanese Writing Basics', 'Get started with Japanese writing systems - Hiragana and Katakana.'),
    ('Russian Language and Culture', 'Explore the Russian language and culture in this introductory course.'),
    ('Arabic for Daily Life', 'Learn Arabic language skills for everyday communication and situations.'),
    ('Portuguese for Business', 'Enhance your Portuguese language skills for professional settings and business contexts.'),
    ('Korean Advanced Grammar', 'Deepen your understanding of advanced Korean grammar concepts and usage.'); -- inserting 10 courses

INSERT INTO lessons (course_id, title, content)
VALUES
    ('d9a9462f-1ba3-4b1a-8f41-cfa13e314d73', 'Introduction to Spanish Alphabet', 'Learn the letters of the Spanish alphabet and their pronunciation.'),
    ('d9a9462f-1ba3-4b1a-8f41-cfa13e314d73', 'Greetings and Introductions', 'Master common greetings and phrases used when meeting people in Spanish-speaking countries.'),
    ('d9a9462f-1ba3-4b1a-8f41-cfa13e314d73', 'Basic Spanish Vocabulary', 'Expand your vocabulary with essential Spanish words and expressions.'),
    ('7a019744-42a1-45ca-965c-92fde402c1e2', 'Intermediate French Conversation', 'Practice conversational French with dialogues and role-playing exercises.'),
    ('7a019744-42a1-45ca-965c-92fde402c1e2', 'French Verb Conjugation', 'Understand verb conjugation rules and patterns in French.'),
    ('b32b1448-925c-4e61-88ab-48e0cbc8541f', 'Noun Cases in German', 'Explore the different noun cases (Nominative, Accusative, Dative, and Genitive) in German.'),
    ('b32b1448-925c-4e61-88ab-48e0cbc8541f', 'German Sentence Structure', 'Learn about word order and sentence structure in German sentences.'),
    ('3b1a6b07-64d4-43f0-937c-55c471ed20d7', 'Introduction to Italian Sounds', 'Understand the sounds of the Italian language and how to pronounce them.'),
    ('3b1a6b07-64d4-43f0-937c-55c471ed20d7', 'Italian Accent Marks', 'Learn about accent marks and their role in Italian pronunciation.'),
    ('12236449-232b-4f9c-a2ae-0be58894b27f', 'Basic Mandarin Chinese Phrases', 'Master basic Mandarin phrases for everyday communication and travel.'),
    ('12236449-232b-4f9c-a2ae-0be58894b27f', 'Introduction to Chinese Characters', 'Learn about Chinese characters (Hanzi) and their structure.'),
    ('9ab5dd15-4e65-4fe7-b8b0-596b8bc105d4', 'Hiragana Basics', 'Learn to read and write Hiragana, one of the Japanese syllabaries.'),
    ('9ab5dd15-4e65-4fe7-b8b0-596b8bc105d4', 'Katakana Basics', 'Learn to read and write Katakana, another Japanese syllabary used for foreign words and loanwords.'),
    ('ff7107bb-b5e5-406b-bdea-4a527a579c70', 'Russian Alphabet (Cyrillic)', 'Get familiar with the Cyrillic alphabet used in the Russian language.'),
    ('ff7107bb-b5e5-406b-bdea-4a527a579c70', 'Russian Pronunciation Rules', 'Understand the pronunciation rules for Russian vowels and consonants.'),
    ('b0f3d272-f76f-493e-be80-4db4ee6b46a9', 'Arabic Greetings and Courtesies', 'Master common Arabic greetings and polite expressions.'),
    ('b0f3d272-f76f-493e-be80-4db4ee6b46a9', 'Arabic Numbers and Counting', 'Learn Arabic numbers and counting up to 100.'),
    ('831cbf71-dc0d-47b6-8fd3-52ffe9c80539', 'Portuguese for Meetings', 'Learn Portuguese phrases and vocabulary for business meetings and discussions.'),
    ('831cbf71-dc0d-47b6-8fd3-52ffe9c80539', 'Portuguese Email Writing', 'Understand the conventions and phrases used in writing professional emails in Portuguese.'),
    ('2796101d-baa6-4c9b-9068-bd2ccea9e85f', 'Advanced Korean Grammar: Particles', 'Explore the various particles used in Korean grammar and their functions.'),
    ('2796101d-baa6-4c9b-9068-bd2ccea9e85f', 'Korean Idioms and Proverbs', 'Learn common Korean idiomatic expressions and proverbs used in everyday speech.'); -- inserting 21 lessons

INSERT INTO enrollments (user_id, course_id, enrollment_date) 
VALUES
('522a5358-6455-4dd8-8a9d-16d8dfc1f23f', '3b1a6b07-64d4-43f0-937c-55c471ed20d7', NOW()), -- Alice Johnson - Italian Pronunciation Mastery
('522a5358-6455-4dd8-8a9d-16d8dfc1f23f', 'b0f3d272-f76f-493e-be80-4db4ee6b46a9', NOW()), -- Alice Johnson - Arabic for Daily Life
('b4d6679a-e39f-43d4-943a-dc07dab4cdab', '2796101d-baa6-4c9b-9068-bd2ccea9e85f', NOW()), -- Bob Williams - Korean Advanced Grammar
('b4d6679a-e39f-43d4-943a-dc07dab4cdab', '7a019744-42a1-45ca-965c-92fde402c1e2', NOW()), -- Bob Williams - French Conversation
('2f864ddf-8efb-4e9b-946d-e08e086ff162', '831cbf71-dc0d-47b6-8fd3-52ffe9c80539', NOW()), -- Charlie Brown - Portuguese for Business
('2f864ddf-8efb-4e9b-946d-e08e086ff162', 'ff7107bb-b5e5-406b-bdea-4a527a579c70', NOW()), -- Charlie Brown - Russian Language and Culture
('66251567-d4c3-4fc1-ab87-b0a91fb491ff', '12236449-232b-4f9c-a2ae-0be58894b27f', NOW()), -- Diana Martinez - Mandarin Chinese for Travelers
('66251567-d4c3-4fc1-ab87-b0a91fb491ff', '9ab5dd15-4e65-4fe7-b8b0-596b8bc105d4', NOW()), -- Diana Martinez - Japanese Writing Basics
('0e54e7ff-f456-452b-8047-fec63cfbc579', 'b0f3d272-f76f-493e-be80-4db4ee6b46a9', NOW()), -- Ella Davis - Arabic for Daily Life
('0e54e7ff-f456-452b-8047-fec63cfbc579', 'b32b1448-925c-4e61-88ab-48e0cbc8541f', NOW()), -- Ella Davis - German Grammar Essentials
('201144db-c584-455d-aad0-6964bfad7d71', 'd9a9462f-1ba3-4b1a-8f41-cfa13e314d73', NOW()), -- Frank Wilson - Spanish for Beginners
('201144db-c584-455d-aad0-6964bfad7d71', '831cbf71-dc0d-47b6-8fd3-52ffe9c80539', NOW()), -- Frank Wilson - Portuguese for Business
('f440f1fb-b8b9-4d5a-ad66-46969c7c217e', '12236449-232b-4f9c-a2ae-0be58894b27f', NOW()), -- Grace Thompson - Mandarin Chinese for Travelers
('f440f1fb-b8b9-4d5a-ad66-46969c7c217e', '3b1a6b07-64d4-43f0-937c-55c471ed20d7', NOW()), -- Grace Thompson - Italian Pronunciation Mastery 
('e9c88fa3-2c9a-468f-9203-a19b92cf3129', 'b32b1448-925c-4e61-88ab-48e0cbc8541f', NOW()), -- Henry Garcia - German Grammar Essentials
('e9c88fa3-2c9a-468f-9203-a19b92cf3129', '7a019744-42a1-45ca-965c-92fde402c1e2', NOW()), -- Henry Garcia - French Conversation
('be92c430-477c-4deb-bde5-2a9840dba615', 'd9a9462f-1ba3-4b1a-8f41-cfa13e314d73', NOW()), -- Ivy Rodriguez - Spanish for Beginners
('be92c430-477c-4deb-bde5-2a9840dba615', '9ab5dd15-4e65-4fe7-b8b0-596b8bc105d4', NOW()), -- Ivy Rodriguez - Japanese Writing Basics
('5602ae7a-70ae-49c9-880f-1590b29880c0', '2796101d-baa6-4c9b-9068-bd2ccea9e85f', NOW()), -- Jack Smith - Korean Advanced Grammar
('5602ae7a-70ae-49c9-880f-1590b29880c0', 'ff7107bb-b5e5-406b-bdea-4a527a579c70', NOW()), -- Jack Smith - Russian Language and Culture
('f440f1fb-b8b9-4d5a-ad66-46969c7c217e', 'b0f3d272-f76f-493e-be80-4db4ee6b46a9', NOW()); -- Grace Thompson - Arabic for Daily Life