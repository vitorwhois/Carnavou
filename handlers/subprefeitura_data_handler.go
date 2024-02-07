package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vitorwhois/Carnavou/models"
)

func BlocosPorDataESubprefeituraHandler(w http.ResponseWriter, r *http.Request) {
	repository := models.ListaRepository{}

	data := r.URL.Query().Get("data")
	subprefeitura := r.URL.Query().Get("subprefeitura")

	if data == "" || subprefeitura == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Data ou Subprefeitura não especificada"))
		return
	}

	blocos, err := repository.ObterBlocosPorDataESubprefeitura(data, subprefeitura)
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
