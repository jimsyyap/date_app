-- Generate fake names using a built-in PostgreSQL function
CREATE FUNCTION generate_random_name() RETURNS text AS $$
BEGIN
    RETURN (SELECT name FROM names ORDER BY RANDOM() LIMIT 1);
END;
$$ LANGUAGE sql;

-- Generate fake emails using a built-in PostgreSQL function
CREATE FUNCTION generate_random_email() RETURNS text AS $$
BEGIN
    RETURN (SELECT email FROM emails ORDER BY RANDOM() LIMIT 1);
END;
$$ LANGUAGE sql;

-- Generate fake passwords using a built-in PostgreSQL function (replace with a stronger hashing algorithm if needed)
CREATE FUNCTION generate_random_password() RETURNS text AS $$
BEGIN
    RETURN (SELECT password FROM passwords ORDER BY RANDOM() LIMIT 1);
END;
$$ LANGUAGE sql;

-- Generate fake avatars using a built-in PostgreSQL function (replace with a more specific approach if needed)
CREATE FUNCTION generate_random_avatar() RETURNS text AS $$
BEGIN
    RETURN 'https://randomuser.me/api/portraits/men/' || (random() * 100)::integer || '.jpg';
END;
$$ LANGUAGE sql;

-- Generate fake dates of birth within a specific range
CREATE FUNCTION generate_random_date_of_birth() RETURNS date AS $$
BEGIN
    RETURN (SELECT date_trunc('year', CURRENT_DATE) - interval '30 years' + interval 'random() * 60 years')::date;
END;
$$ LANGUAGE sql;

-- Generate fake occupations from a predefined list
CREATE FUNCTION generate_random_occupation() RETURNS text AS $$
BEGIN
    RETURN (SELECT occupation FROM occupations ORDER BY RANDOM() LIMIT 1);
END;
$$ LANGUAGE sql;

-- Generate fake user preferences (adjust as needed)
CREATE FUNCTION generate_random_preferences() RETURNS user_preferences AS $$
BEGIN
    RETURN (SELECT
        generate_random_integer(18, 35) AS min_age,
        generate_random_integer(36, 55) AS max_age,
        generate_random_integer(1, 100) AS max_distance,
        (SELECT gender FROM genders ORDER BY RANDOM() LIMIT 1) AS gender_preference
    );
END;
$$ LANGUAGE sql;

-- Insert 100 fake users
INSERT INTO users (name, email, password_hash, avatar_url, date_of_birth, occupation)
SELECT
    generate_random_name(),
    generate_random_email(),
    generate_random_password(),
    generate_random_avatar(),
    generate_random_date_of_birth(),
    generate_random_occupation()
FROM generate_series(1, 100);

-- Insert 100 fake conversations (assuming random pairings)
INSERT INTO conversations (user1_id, user2_id)
SELECT
    user1.id,
    user2.id
FROM users user1
JOIN users user2 ON user1.id < user2.id
ORDER BY RANDOM()
LIMIT 100;

-- Insert 100 fake messages (assuming random assignments to conversations)
INSERT INTO messages (conversation_id, sender_id, content)
SELECT
    (SELECT id FROM conversations ORDER BY RANDOM() LIMIT 1),
    (SELECT id FROM users ORDER BY RANDOM() LIMIT 1),
    'This is a fake message.'
FROM generate_series(1, 100);

-- Insert 100 fake matches (assuming random pairings)
INSERT INTO matches (user1_id, user2_id, status)
SELECT
    user1.id,
    user2.id,
    'pending'
FROM users user1
JOIN users user2 ON user1.id < user2.id
ORDER BY RANDOM()
LIMIT 100;

-- Insert 100 fake user preferences
INSERT INTO user_preferences (user_id)
SELECT id FROM users;

-- Update user preferences with random values
UPDATE user_preferences
SET min_age = generate_random_preferences().min_age,
    max_age = generate_random_preferences().max_age,
    max_distance = generate_random_preferences().max_distance,
    gender_preference = generate_random_preferences().gender_preference;
