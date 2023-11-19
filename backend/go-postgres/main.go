package main

import (
	"database/sql"
	"github.com/arioprima/Point-Of-Sale/tree/main/backend/go-postgres/initializers"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Println("🚀 Could not load environment variables", err)
		return
	}

	db, err := initializers.ConnectDB(&config)
	if err != nil {
		log.Println("🚀 Could not connect to database", err)
		return
	}

	// Check if db is connected
	err = db.Ping()
	if err != nil {
		log.Println("🚀 Could not ping database", err)
		return
	}

	log.Println("🚀 Database connection successful")

	// Close the database connection when done
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("🚀 Could not close database connection", err)
			return
		}
	}(db)

}
