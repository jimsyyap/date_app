import psycopg2
from faker import Faker
from datetime import datetime, timedelta
import random

# Database connection parameters
DB_PARAMS = {
    "dbname": "dateapp",
    "user": "jim",
    "password": "whatsimportantnow",
    "host": "localhost",
    "port": "5432"
}

fake = Faker()

def create_connection():
    return psycopg2.connect(**DB_PARAMS)

def generate_users(num_users=100):
    users = []
    for _ in range(num_users):
        user = (
            fake.name(),
            fake.email(),
            fake.password(),
            fake.image_url(),
            fake.date_of_birth(minimum_age=18, maximum_age=80),
            fake.job(),
        )
        users.append(user)
    return users

def insert_users(conn, users):
    with conn.cursor() as cur:
        cur.executemany(
            """
            INSERT INTO users (name, email, password_hash, avatar_url, date_of_birth, occupation)
            VALUES (%s, %s, %s, %s, %s, %s)
            RETURNING id
            """,
            users
        )
        return [row[0] for row in cur.fetchall()]

def generate_conversations(user_ids, num_conversations=50):
    conversations = []
    for _ in range(num_conversations):
        user1, user2 = random.sample(user_ids, 2)
        conversations.append((user1, user2))
    return conversations

def insert_conversations(conn, conversations):
    with conn.cursor() as cur:
        cur.executemany(
            """
            INSERT INTO conversations (user1_id, user2_id)
            VALUES (%s, %s)
            RETURNING id
            """,
            conversations
        )
        return [row[0] for row in cur.fetchall()]

def generate_messages(conversation_ids, user_ids, num_messages=200):
    messages = []
    for _ in range(num_messages):
        conversation_id = random.choice(conversation_ids)
        sender_id = random.choice(user_ids)
        messages.append((
            conversation_id,
            sender_id,
            fake.text(max_nb_chars=200),
            fake.date_time_between(start_date="-30d", end_date="now")
        ))
    return messages

def insert_messages(conn, messages):
    with conn.cursor() as cur:
        cur.executemany(
            """
            INSERT INTO messages (conversation_id, sender_id, content, sent_at)
            VALUES (%s, %s, %s, %s)
            """,
            messages
        )

def generate_matches(user_ids, num_matches=150):
    matches = []
    for _ in range(num_matches):
        user1, user2 = random.sample(user_ids, 2)
        status = random.choice(['pending', 'accepted', 'rejected'])
        matches.append((user1, user2, status))
    return matches

def insert_matches(conn, matches):
    with conn.cursor() as cur:
        cur.executemany(
            """
            INSERT INTO matches (user1_id, user2_id, status)
            VALUES (%s, %s, %s)
            """,
            matches
        )

def generate_user_preferences(user_ids):
    preferences = []
    for user_id in user_ids:
        preferences.append((
            user_id,
            random.randint(18, 50),
            random.randint(25, 80),
            random.randint(5, 100),
            random.choice(['male', 'female', 'any'])
        ))
    return preferences

def insert_user_preferences(conn, preferences):
    with conn.cursor() as cur:
        cur.executemany(
            """
            INSERT INTO user_preferences (user_id, min_age, max_age, max_distance, gender_preference)
            VALUES (%s, %s, %s, %s, %s)
            """,
            preferences
        )

def main():
    conn = create_connection()
    try:
        # Generate and insert users
        users = generate_users(100)
        user_ids = insert_users(conn, users)

        # Generate and insert conversations
        conversations = generate_conversations(user_ids, 50)
        conversation_ids = insert_conversations(conn, conversations)

        # Generate and insert messages
        messages = generate_messages(conversation_ids, user_ids, 200)
        insert_messages(conn, messages)

        # Generate and insert matches
        matches = generate_matches(user_ids, 150)
        insert_matches(conn, matches)

        # Generate and insert user preferences
        preferences = generate_user_preferences(user_ids)
        insert_user_preferences(conn, preferences)

        conn.commit()
        print("Fake data generation complete!")
    except Exception as e:
        conn.rollback()
        print(f"An error occurred: {e}")
    finally:
        conn.close()

if __name__ == "__main__":
    main()
