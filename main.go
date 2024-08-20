package main

import (
	"database/sql"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

//go:embed frontend/build/*
var static embed.FS

const (
	host     = "localhost"
	port     = 5432
	user     = "jim"
	password = "whatsimportantnow"
	dbname   = "dateapp"
)

func main() {
	// Set up database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
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

	fmt.Println("Successfully connected to database!")

	// Set up file server for frontend
	fsys, err := fs.Sub(static, "frontend/build")
	if err != nil {
		log.Fatal(err)
	}

	// Set up API routes
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name FROM your_table LIMIT 10")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var results []map[string]interface{}
		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			results = append(results, map[string]interface{}{
				"id":   id,
				"name": name,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	})

	// Serve static files
	http.Handle("/", http.FileServer(http.FS(fsys)))

	// Start the server
	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
