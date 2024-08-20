package main

import (
	"database/sql"
	"embed"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/fs"
	"log"
	"net/http"
)

//go:embed frontend/build

var static embed.FS

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)
