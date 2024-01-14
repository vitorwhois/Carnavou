package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vitorwhois/Carnavou/models"
)

// BlocosPorDataHandler manipula solicitações para obter blocos por data
func BlocosPorDataHandler(w http.ResponseWriter, r *http.Request) {
	// Lógica para obter dados de blocos por data usando o serviço e repositório correspondentes
	repository := models.ListaRepository{}

	// Recupera a data da solicitação HTTP atual
	data := r.URL.Query().Get("data")
	if data == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Data não especificada"))
		return
	}

	blocosData, err := repository.ObterBlocosPorData(data)
	if err != nil {
		log.Printf("Erro ao obter dados de blocos por data: %v", err)
		http.Error(w, "Erro ao obter dados de blocos por data", http.StatusInternalServerError)
		return
	}

	// Converte os dados para JSON
	jsonData, err := json.Marshal(blocosData)
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
