package initializers

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectDB(config *Config) (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("Connect to database error", err)
	}
	err = db.Ping()
	if err != nil {
		log.Println("Ping to database error", err)
	} else {
		log.Println("Connect to database successfully")
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db, err
}
