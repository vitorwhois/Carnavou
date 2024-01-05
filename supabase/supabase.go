package supabase

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

// SupabaseDB é uma instância do banco de dados Supabase
var SupabaseDB *sql.DB

func init() {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar as variáveis de ambiente:", err)
	}
	// Atribui os valores das variáveis de ambiente
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
