package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var SupabaseDB *sql.DB

func init() {

	env := os.Getenv("ENV")

	var user, password, host, port, dbname string

	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Erro ao carregar .env file")
		}

		user = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		dbname = os.Getenv("DB_NAME")

	}

	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", user, password, host, port, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados:", err)
	}
	SupabaseDB = db
}

func OpenDatabaseConnection() (*sql.DB, error) {
	return SupabaseDB, nil
}
