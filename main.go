package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed frontend/build/*
var static embed.FS

func main() {
	fsys, err := fs.Sub(static, "frontend/build")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(fsys)))

	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
