package models

import (
	"database/sql"
	"log"

	"github.com/vitorwhois/Carnavou/supabase"
)

// ListaRepository define um repositório para operações relacionadas à lista (blocos, no seu caso)
type ListaRepository struct{}

// ObterBlocosPorData retorna blocos com base na data fornecida
func (lr *ListaRepository) ObterBlocosPorData(data string) ([]Bloco, error) {
	// Realiza a consulta na tabela "blocos" para as entradas com a data fornecida
	rows, err := supabase.SupabaseDB.Query("SELECT id, nome, data, subprefeitura, local, concentracao, tamanho_id FROM blocos WHERE data = $1", data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	// Processa os resultados
	var blocos []Bloco
	for rows.Next() {
		var bloco Bloco
		var tamanho sql.NullString

		err := rows.Scan(&bloco.ID, &bloco.Nome, &bloco.Data, &bloco.Subprefeitura, &bloco.Local, &bloco.Concentracao, &tamanho)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		// Verifica se há um tamanho válido
		if tamanho.Valid {
			bloco.Tamanho = tamanho.String
		} else {
			bloco.Tamanho = "NULL"
		}

		blocos = append(blocos, bloco)
	}

	// Verifica se houve algum erro durante o processamento das linhas
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return blocos, nil
}
