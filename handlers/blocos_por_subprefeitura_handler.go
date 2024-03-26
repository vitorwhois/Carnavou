package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/vitorwhois/Carnavou/repositories"
)

type BlocosPorSubprefeituraHandler struct {
	repo *repositories.BlocosPorSubprefeituraRepository
}

func NewBlocosPorSubprefeituraHandler(db *sql.DB) *BlocosPorSubprefeituraHandler {
	repo := repositories.NewBlocosPorSubprefeituraRepository(db)
	return &BlocosPorSubprefeituraHandler{repo: repo}
}

func (h *BlocosPorSubprefeituraHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Recupera o subprefeitura da solicitação HTTP atual
	subprefeitura := r.URL.Query().Get("subprefeituras")
	if subprefeitura == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Subprefeitura não especificada"))
		return
	}

	blocosSubprefeitura, err := h.repo.ObterBlocosPorSubprefeitura(subprefeitura)
	if err != nil {
		log.Printf("Erro ao obter dados de blocos por subprefeitura: %v", err)
		http.Error(w, "Erro ao obter dados de blocos por subprefeitura", http.StatusInternalServerError)
		return
	}

	// Converte os dados para JSON
	jsonData, err := json.Marshal(blocosSubprefeitura)
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
