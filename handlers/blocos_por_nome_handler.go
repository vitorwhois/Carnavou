package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/vitorwhois/Carnavou/repositories"
)

type BlocosPorNomeHandler struct {
	repo *repositories.BlocoPorNomeRepository
}

func NewBlocosPorNomeHandler(db *sql.DB) *BlocosPorNomeHandler {
	repo := repositories.NewBlocoPorNomeRepository(db)
	return &BlocosPorNomeHandler{repo: repo}
}

func (h *BlocosPorNomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Recupera o nome da solicitação HTTP atual
	nome := r.URL.Query().Get("nomes")
	if nome == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Nome não especificado"))
		return
	}

	blocosNome, err := h.repo.ObterBlocosPorNome(nome)
	if err != nil {
		log.Printf("Erro ao obter dados de blocos por nome: %v", err)
		http.Error(w, "Erro ao obter dados de blocos por nome", http.StatusInternalServerError)
		return
	}

	// Converte os dados para JSON
	jsonData, err := json.Marshal(blocosNome)
	if err != nil {
		log.Printf("Erro ao converter dados para JSON: %v", err)
		http.Error(w, "Erro ao converter dados para JSON", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho Content-Type para application/json
	w.Header().Set("Content-Type", "application/json")

	// Responde com os dados em formato JSON
	w.Write(jsonData)
}
