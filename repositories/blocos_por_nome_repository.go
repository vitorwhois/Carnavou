package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vitorwhois/Carnavou/models"
)

type BlocoPorNomeRepository struct {
	db *sql.DB
}

func NewBlocoPorNomeRepository(db *sql.DB) *BlocoPorNomeRepository {
	return &BlocoPorNomeRepository{db: db}
}

func (br *BlocoPorNomeRepository) ObterBlocosPorNome(nome string) ([]models.Bloco, error) {
	rows, err := br.db.Query("SELECT id, nome, data, subprefeitura, local, concentracao, tamanho_id FROM blocos WHERE nome ILIKE '%' || $1 || '%' ORDER BY nome, data ", nome)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar a consulta: %v", err)
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
