package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type PostgresDB struct {
	db *sqlx.DB
}

// GetDB creates a new PostgresDB instance using the provided secrets and returns it.
func GetDB(secrets map[string]string) *PostgresDB {
	db := createDB(secrets)
	return &PostgresDB{
		db: db,
	}
}

func createDB(secrets map[string]string) *sqlx.DB {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file", err)
		return nil
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv(("DB_PORT"))

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s",
		dbHost, dbPort, secrets["username"], secrets["password"])

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db

}
