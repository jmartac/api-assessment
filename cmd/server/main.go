package main

import (
	"api-assessment/internal"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const (
	portDefault = 3000
)

func main() {
	var port int

	// Parse command line port flag
	flag.IntVar(&port, "p", portDefault, "API port")
	flag.Parse()

	services, err := internal.NewServices()
	if err != nil {
		panic(err)
	}

	// Close database connection when main() exits
	defer func(db *internal.Services) {
		log.Println("Disconnecting from database...")
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(services)

	// Migrate database on startup
	log.Println("Migrating database...")
	err = services.AutoMigrate()
	if err != nil {
		panic(err)
	}

	router := internal.NewRouter(services)

	log.Printf("Starting server on port %d...", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
