package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
	"github.com/vitorwhois/Carnavou/models"
)

type BlocoPorIDRepository struct {
	db *sql.DB
}

func NewBlocoPorIDRepository(db *sql.DB) *BlocoPorIDRepository {
	return &BlocoPorIDRepository{db: db}
}

func (r *BlocoPorIDRepository) ObterBlocosPorIDs(ids []string) ([]models.Bloco, error) {
	query := fmt.Sprintf("SELECT * FROM blocos WHERE id IN (%s)", strings.Join(ids, ","))
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("falha ao executar a consulta: %w", err)
	}
	defer rows.Close()

	var blocos []models.Bloco
	for rows.Next() {
		var bloco models.Bloco
		var tamanho sql.NullString

		err := rows.Scan(&bloco.ID, &bloco.Nome, &bloco.Data, &bloco.Subprefeitura, &bloco.Local, &bloco.Concentracao, &tamanho)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		if tamanho.Valid {
			bloco.Tamanho = tamanho.String
		} else {
			bloco.Tamanho = "NULL"
		}

		blocos = append(blocos, bloco)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return blocos, nil
}
