-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    avatar_url VARCHAR(255),
    date_of_birth DATE NOT NULL,
    occupation VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Conversations table
CREATE TABLE conversations (
    id SERIAL PRIMARY KEY,
    user1_id INTEGER REFERENCES users(id),
    user2_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Messages table
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    conversation_id INTEGER REFERENCES conversations(id),
    sender_id INTEGER REFERENCES users(id),
    content TEXT NOT NULL,
    sent_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Matches table
CREATE TABLE matches (
    id SERIAL PRIMARY KEY,
    user1_id INTEGER REFERENCES users(id),
    user2_id INTEGER REFERENCES users(id),
    status VARCHAR(50) NOT NULL, -- e.g., 'pending', 'accepted', 'rejected'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- User preferences table (for filtering)
CREATE TABLE user_preferences (
    user_id INTEGER PRIMARY KEY REFERENCES users(id),
    min_age INTEGER,
    max_age INTEGER,
    max_distance INTEGER,
    gender_preference VARCHAR(50),
    -- Add more preference fields as needed
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
