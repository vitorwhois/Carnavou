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
	if env == "" {
		log.Println("Variável de ambiente ENV não configurada. Usando valor padrão 'development'.")
		env = "development"
	}
	log.Println("Ambiente atual:", env)

	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Erro ao carregar o arquivo .env:", err)
		}
		log.Println(".env carregado com sucesso")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	if user == "" || password == "" || host == "" || port == "" || dbname == "" {
		log.Fatal("Alguma variável de ambiente do banco de dados está faltando.")
	}

	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user, password, host, port, dbname,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados:", err)
	}
	log.Println("Conexão com o banco de dados estabelecida com sucesso")

	SupabaseDB = db
}

// Função para retornar a conexão com o banco
func OpenDatabaseConnection() (*sql.DB, error) {
	return SupabaseDB, nil
}
