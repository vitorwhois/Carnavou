package storage

import (
	"database/sql"
	"log"
	"os"

	//	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var SupabaseDB *sql.DB

func init() {
	//	err := godotenv.Load()
	//	if err != nil {
	//		log.Fatal("Erro ao carregar .env file")
	//	}

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

/* conectar local
package storage
import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
var (
	supabaseUrl, supabaseKey, postgresHost, postgresPassword string
)
var SupabaseDB *sql.DB

func init() {
	// Retirar para fazer o deploy no Render

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar as variáveis de ambiente:", err)
	}
	supabaseUrl = os.Getenv("SUPABASE_URL")
	supabaseKey = os.Getenv("SUPABASE_KEY")
	postgresHost = os.Getenv("POSTGRES_HOST")
	postgresPassword = os.Getenv("POSTGRES_PASSWORD")
	db, err := sql.Open("postgres", fmt.Sprintf("user=postgres password=%s host=%s port=5432 dbname=postgres", postgresPassword, postgresHost))
	if err != nil {
		log.Fatal(err)
	}
	SupabaseDB = db
}
func OpenDatabaseConnection() (*sql.DB, error) {
	// Função para abrir a conexão com o banco de dados
	// Retornará um ponteiro para o DB para ser usado em outras partes do código
	return SupabaseDB, nil
}
*/
