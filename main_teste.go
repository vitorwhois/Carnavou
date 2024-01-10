package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/vitorwhois/Carnavou/handlers"
	"github.com/vitorwhois/Carnavou/models"
	"github.com/vitorwhois/Carnavou/storage"
)

func main() {
	// Crie uma instância do IndexHandler
	indexHandler := handlers.NewIndexHandler()

	// Use o método Handle da instância
	http.HandleFunc("/", indexHandler.Handle)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/blocosPorData", BlocosPorDataHandler) // Ajuste da rota para /blocosPorData

	//Abre conexão com o banco de dados
	db, err := storage.OpenDatabaseConnection()
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados", err)
	}
	defer db.Close() //Fechar a conexão com o banco de dados

	port := 8080
	fmt.Printf("Servidor rodando em http://localhost:%d\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}

}

// BlocosPorDataHandler manipula solicitações para obter blocos por data
func BlocosPorDataHandler(w http.ResponseWriter, r *http.Request) {
	// Lógica para obter dados de blocos por data usando o serviço e repositório correspondentes
	repository := models.ListaRepository{}
	blocosData, err := repository.ObterBlocosPorData("17/02/2024") // Data desejada
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
