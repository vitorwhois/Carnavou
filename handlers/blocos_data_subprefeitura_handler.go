package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/vitorwhois/Carnavou/repositories"
)

type BlocosPorDataSubprefeituraHandler struct {
	repo *repositories.BlocosPorDataSubprefeituraRepository
}

func NewBlocosPorDataSubprefeituraHandler(db *sql.DB) *BlocosPorDataSubprefeituraHandler {
	repo := repositories.NewBlocosPorDataSubprefeituraRepository(db)
	return &BlocosPorDataSubprefeituraHandler{repo: repo}
}

func (h *BlocosPorDataSubprefeituraHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("datas")
	subprefeitura := r.URL.Query().Get("subprefeituras")

	if data == "" || subprefeitura == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Data ou Subprefeitura não especificada"))
		return
	}

	blocos, err := h.repo.ObterBlocosPorDataSubprefeitura(data, subprefeitura)
	if err != nil {
		log.Printf("Erro ao obter dados de blocos por data e subprefeitura: %v", err)
		http.Error(w, "Erro ao obter dados de blocos por data e subprefeitura", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(blocos)
	if err != nil {
		log.Printf("Erro ao converter dados para JSON: %v", err)
		http.Error(w, "Erro ao converter dados para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
