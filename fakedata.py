import psycopg2
from faker import Faker
from datetime import datetime, timedelta
import random
import sys

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
    try:
        conn = psycopg2.connect(**DB_PARAMS)
        print("Successfully connected to the database.")
        return conn
    except psycopg2.Error as e:
        print(f"Unable to connect to the database: {e}")
        sys.exit(1)

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
        inserted_ids = []
        for user in users:
            try:
                cur.execute(
                    """
                    INSERT INTO users (name, email, password_hash, avatar_url, date_of_birth, occupation)
                    VALUES (%s, %s, %s, %s, %s, %s)
                    RETURNING id
                    """,
                    user
                )
                result = cur.fetchone()
                if result:
                    inserted_ids.append(result[0])
                else:
                    print(f"Warning: No ID returned for user {user[1]}")
            except psycopg2.Error as e:
                print(f"Error inserting user {user[1]}: {e}")
        return inserted_ids

def insert_conversations(conn, conversations):
    with conn.cursor() as cur:
        inserted_ids = []
        for conversation in conversations:
            try:
                cur.execute(
                    """
                    INSERT INTO conversations (user1_id, user2_id)
                    VALUES (%s, %s)
                    RETURNING id
                    """,
                    conversation
                )
                result = cur.fetchone()
                if result:
                    inserted_ids.append(result[0])
                else:
                    print(f"Warning: No ID returned for conversation {conversation}")
            except psycopg2.Error as e:
                print(f"Error inserting conversation: {e}")
        return inserted_ids

# ... (keep other functions as they are)

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

def main():
    conn = create_connection()
    try:
        # Generate and insert users
        users = generate_users(100)
        user_ids = insert_users(conn, users)
        if not user_ids:
            raise Exception("No users were inserted successfully.")
        print(f"Inserted {len(user_ids)} users.")

        # Generate and insert conversations
        conversations = generate_conversations(user_ids, 50)
        conversation_ids = insert_conversations(conn, conversations)
        if not conversation_ids:
            raise Exception("No conversations were inserted successfully.")
        print(f"Inserted {len(conversation_ids)} conversations.")

        # Generate and insert messages
        messages = generate_messages(conversation_ids, user_ids, 200)
        insert_messages(conn, messages)
        print("Inserted messages.")

        # Generate and insert matches
        matches = generate_matches(user_ids, 150)
        insert_matches(conn, matches)
        print("Inserted matches.")

        # Generate and insert user preferences
        preferences = generate_user_preferences(user_ids)
        insert_user_preferences(conn, preferences)
        print("Inserted user preferences.")

        conn.commit()
        print("Fake data generation complete!")
    except Exception as e:
        conn.rollback()
        print(f"An error occurred: {e}")
        print("Traceback:", sys.exc_info())
    finally:
        conn.close()

if __name__ == "__main__":
    main()
