package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "jim"
	password = "whatsimportantnow"
	dbname   = "dateapp"
)

func main() {
	// Create a new random source and generator
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// Connect to the database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")

	// Generate dummy data
	generateUsers(db, rng, 100)
	generateConversations(db, rng, 50)
	generateMessages(db, rng, 200)
	generateMatches(db, rng, 30)
	generateUserPreferences(db, rng, 100)

	fmt.Println("Dummy data generation completed")
}

func generateUsers(db *sql.DB, rng *rand.Rand, count int) {
	for i := 0; i < count; i++ {
		_, err := db.Exec(`
			INSERT INTO users (name, email, password_hash, avatar_url, date_of_birth, occupation)
			VALUES ($1, $2, $3, $4, $5, $6)
		`, faker.Name(), faker.Email(), faker.Password(), faker.URL(), faker.Date(), faker.Word())

		if err != nil {
			log.Printf("Error inserting user: %v", err)
		}
	}
	fmt.Printf("Generated %d users\n", count)
}

func generateConversations(db *sql.DB, rng *rand.Rand, count int) {
	for i := 0; i < count; i++ {
		_, err := db.Exec(`
			INSERT INTO conversations (user1_id, user2_id)
			VALUES ($1, $2)
		`, rng.Intn(100)+1, rng.Intn(100)+1)

		if err != nil {
			log.Printf("Error inserting conversation: %v", err)
		}
	}
	fmt.Printf("Generated %d conversations\n", count)
}

func generateMessages(db *sql.DB, rng *rand.Rand, count int) {
	for i := 0; i < count; i++ {
		_, err := db.Exec(`
			INSERT INTO messages (conversation_id, sender_id, content)
			VALUES ($1, $2, $3)
		`, rng.Intn(50)+1, rng.Intn(100)+1, faker.Sentence())

		if err != nil {
			log.Printf("Error inserting message: %v", err)
		}
	}
	fmt.Printf("Generated %d messages\n", count)
}

func generateMatches(db *sql.DB, rng *rand.Rand, count int) {
	statuses := []string{"pending", "accepted", "rejected"}
	for i := 0; i < count; i++ {
		_, err := db.Exec(`
			INSERT INTO matches (user1_id, user2_id, status)
			VALUES ($1, $2, $3)
		`, rng.Intn(100)+1, rng.Intn(100)+1, statuses[rng.Intn(len(statuses))])

		if err != nil {
			log.Printf("Error inserting match: %v", err)
		}
	}
	fmt.Printf("Generated %d matches\n", count)
}

func generateUserPreferences(db *sql.DB, rng *rand.Rand, count int) {
	genders := []string{"male", "female", "non-binary", "any"}
	for i := 0; i < count; i++ {
		_, err := db.Exec(`
			INSERT INTO user_preferences (user_id, min_age, max_age, max_distance, gender_preference)
			VALUES ($1, $2, $3, $4, $5)
		`, i+1, rng.Intn(10)+18, rng.Intn(20)+30, rng.Intn(50)+10, genders[rng.Intn(len(genders))])

		if err != nil {
			log.Printf("Error inserting user preference: %v", err)
		}
	}
	fmt.Printf("Generated %d user preferences\n", count)
}
