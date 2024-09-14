package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/lib/pq"
)

// Replace with your database connection string
const connStr = "postgres://user:password@localhost:5432/your_database"

func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Generate and insert dummy data
	generateAndInsertUsers(db, 100)
	generateAndInsertConversations(db)
	generateAndInsertMessages(db)
	generateAndInsertMatches(db)
	generateAndInsertUserPreferences(db)
}

func generateAndInsertUsers(db *sql.DB, numUsers int) {
	faker.Seed(time.Now().UnixNano())
	for i := 0; i < numUsers; i++ {
		name := faker.Name().FirstName()
		email := faker.Internet().Email()
		passwordHash := faker.Internet().Password(10)
		avatarURL := faker.Internet().URL("image")
		dateOfBirth := faker.Date().Birthday()
		occupation := faker.Job().Title()

		_, err := db.Exec("INSERT INTO users (name, email, password_hash, avatar_url, date_of_birth, occupation) VALUES ($1, $2, $3, $4, $5, $6)",
			name, email, passwordHash, avatarURL, dateOfBirth, occupation)
		if err != nil {
			panic(err)
		}
	}
}

func generateAndInsertConversations(db *sql.DB, numConversations int) {
	faker.Seed(time.Now().UnixNano())
	for i := 0; i < numConversations; i++ {
		user1ID := rand.Intn(100) + 1 // Assuming 100 users
		user2ID := rand.Intn(100) + 1
		if user1ID == user2ID {
			user2ID = (user1ID+1)%100 + 1
		}

		_, err := db.Exec("INSERT INTO conversations (user1_id, user2_id) VALUES ($1, $2)", user1ID, user2ID)
		if err != nil {
			panic(err)
		}
	}
}

func generateAndInsertMessages(db *sql.DB, numMessages int) {
	faker.Seed(time.Now().UnixNano())
	for i := 0; i < numMessages; i++ {
		conversationID := rand.Intn(100) + 1 // Assuming 100 conversations
		senderID := rand.Intn(100) + 1
		content := faker.Lorem().Sentence()

		_, err := db.Exec("INSERT INTO messages (conversation_id, sender_id, content) VALUES ($1, $2, $3)", conversationID, senderID, content)
		if err != nil {
			panic(err)
		}
	}
}

func generateAndInsertMatches(db *sql.DB, numMatches int) {
	faker.Seed(time.Now().UnixNano())
	for i := 0; i < numMatches; i++ {
		user1ID := rand.Intn(100) + 1 // Assuming 100 users
		user2ID := rand.Intn(100) + 1
		if user1ID == user2ID {
			user2ID = (user1ID+1)%100 + 1
		}

		status := faker.Random().Element([]string{"pending", "accepted", "rejected"})

		_, err := db.Exec("INSERT INTO matches (user1_id, user2_id, status) VALUES ($1, $2, $3)", user1ID, user2ID, status)
		if err != nil {
			panic(err)
		}
	}
}

func generateAndInsertUserPreferences(db *sql.DB, numPreferences int) {
	faker.Seed(time.Now().UnixNano())
	for i := 0; i < numPreferences; i++ {
		userID := rand.Intn(100) + 1 // Assuming 100 users
		minAge := rand.Intn(50) + 18
		maxAge := minAge + rand.Intn(30) + 1
		maxDistance := rand.Intn(100) + 1
		genderPreference := faker.Random().Element([]string{"male", "female", "other"})

		_, err := db.Exec("INSERT INTO user_preferences (user_id, min_age, max_age, max_distance, gender_preference) VALUES ($1, $2, $3, $4, $5)",
			userID, minAge, maxAge, maxDistance, genderPreference)
		if err != nil {
			panic(err)
		}
	}
}
