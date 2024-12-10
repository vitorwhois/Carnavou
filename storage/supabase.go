package storage

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var SupabaseDB *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	connectionString := "user=" + user + " password=" + password + " host=" + host + " port=" + port + " dbname=" + dbname

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	SupabaseDB = db
}

func OpenDatabaseConnection() (*sql.DB, error) {
	return SupabaseDB, nil
}
