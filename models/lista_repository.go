package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vitorwhois/Carnavou/storage"
)

type ListaRepository struct{}

func (lr *ListaRepository) ObterBlocosPorData(data string) ([]Bloco, error) {
	db, err := storage.OpenDatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("Erro ao obter a conexão com o banco de dados: %v", err)
	}

	rows, err := db.Query("SELECT id, nome, data, subprefeitura, local, concentracao, tamanho_id FROM blocos WHERE data = $1", data)
	if err != nil {
		return nil, fmt.Errorf("Erro ao executar a consulta: %v", err)
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
