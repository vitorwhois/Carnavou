package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/vitorwhois/Carnavou/repositories"
)

type BlocosPorIDHandler struct {
	repo *repositories.BlocoPorIDRepository
}

func NewBlocosPorIDHandler(db *sql.DB) *BlocosPorIDHandler {
	repo := repositories.NewBlocoPorIDRepository(db)
	return &BlocosPorIDHandler{repo: repo}
}

func (h *BlocosPorIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ids := strings.Split(r.URL.Query().Get("ids"), ",")
	if len(ids) == 0 {
		http.Error(w, "IDs de blocos não fornecidos", http.StatusBadRequest)
		return
	}

	blocos, err := h.repo.ObterBlocosPorIDs(ids)
	if err != nil {
		log.Printf("Erro ao obter dados de blocos por ID: %v", err)
		http.Error(w, "Erro ao obter dados de blocos por ID", http.StatusInternalServerError)
		return
	}

	if len(blocos) == 0 {
		http.Error(w, "Nenhum bloco encontrado com os IDs fornecidos", http.StatusNotFound)
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
