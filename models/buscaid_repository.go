package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
	"github.com/vitorwhois/Carnavou/storage"
)

type ObterBlocosPorIDs struct{}

func (r *ListaRepository) ObterBlocosPorIDs(ids []string) ([]Bloco, error) {
	db, err := storage.OpenDatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("Erro ao obter a conexão com o banco de dados: %v", err)
	}

	query := fmt.Sprintf("SELECT * FROM blocos WHERE id IN (%s)", strings.Join(ids, ","))
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Falha ao executar a consulta: %w", err)
	}
	defer rows.Close()

	var blocos []Bloco
	for rows.Next() {
		var bloco Bloco
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
