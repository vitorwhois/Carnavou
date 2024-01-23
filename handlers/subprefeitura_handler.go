package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vitorwhois/Carnavou/models"
)

// PesquisarBlocosHandler manipula solicitações para pesquisar blocos por local
func SubprefeituraBlocosHandler(w http.ResponseWriter, r *http.Request) {
	// Lógica para pesquisar dados de blocos por subprefeitura usando o serviço e repositório correspondentes
	repository := models.SubprefeituraRepository{}

	// Recupera o subprefeitura da solicitação HTTP atual
	subprefeitura := r.URL.Query().Get("Subprefeitura")
	if subprefeitura == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Subprefeitura não especificada"))
		return
	}

	blocosSubprefeitura, err := repository.ObterBlocosPorSubprefeitura(subprefeitura)
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
