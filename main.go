package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/vitorwhois/Carnavou/handlers"
	"github.com/vitorwhois/Carnavou/storage"
)

func main() {
	// Crie uma instância do IndexHandler
	indexHandler := handlers.NewIndexHandler()

	// Use o método Handle da instância
	http.HandleFunc("/", indexHandler.Handle)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/blocosPorData", handlers.BlocosPorDataHandler)     // Ajuste da rota para /blocosPorData
	http.HandleFunc("/pesquisarBlocos", handlers.PesquisarBlocosHandler) // Pesquisa por nome

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
