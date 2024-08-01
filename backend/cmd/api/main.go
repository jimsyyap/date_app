package main

import (
	"flag"
	"fmt"
	"net/http"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string
	Version     string
	Environment string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "app environment (development|production)")
	flag.Parse()

	fmt.Println("port:", cfg.port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
