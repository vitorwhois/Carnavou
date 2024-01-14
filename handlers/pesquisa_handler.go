package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vitorwhois/Carnavou/models"
)

// PesquisarBlocosHandler manipula solicitações para pesquisar blocos por nome
func PesquisarBlocosHandler(w http.ResponseWriter, r *http.Request) {
	// Lógica para pesquisar dados de blocos por nome usando o serviço e repositório correspondentes
	repository := models.BlocoRepository{}

	// Recupera o nome da solicitação HTTP atual
	nome := r.URL.Query().Get("nome")
	if nome == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Nome não especificado"))
		return
	}

	blocosNome, err := repository.ObterBlocosPorNome(nome)
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

/* // PesquisarBlocosHandler manipula solicitações para pesquisar blocos por nome
func PesquisarBlocosHandler(w http.ResponseWriter, r *http.Request) {
	// Obtenha o termo de pesquisa da consulta
	termoPesquisa := r.FormValue("nomeInput")

	if termoPesquisa == "" {
		http.Error(w, "Termo de pesquisa vazio", http.StatusBadRequest)
		return
	}

	// Lógica para obter dados de blocos por nome usando o serviço e repositório correspondentes
	repository := models.BlocoRepository{}
	blocos, err := repository.ObterBlocosPorNome(termoPesquisa)
	if err != nil {
		log.Printf("Erro ao obter dados de blocos por nome: %v", err)
		// Modificar para responder com JSON mesmo em caso de erro
		respondWithError(w, "Erro ao obter dados de blocos por nome", http.StatusInternalServerError)
		return
	}

	// Converte os dados para JSON
	jsonData, err := json.Marshal(blocos)
	if err != nil {
		log.Printf("Erro ao converter dados para JSON: %v", err)
		// Modificar para responder com JSON mesmo em caso de erro
		respondWithError(w, "Erro ao converter dados para JSON", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho Content-Type para application/json
	w.Header().Set("Content-Type", "application/json")

	// Responde com os dados em formato JSON
	w.Write(jsonData)
}

// Função auxiliar para responder com erro em formato JSON
func respondWithError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	// Responde com uma mensagem de erro em formato JSON
	json.NewEncoder(w).Encode(map[string]interface{}{"error": message})
} */
