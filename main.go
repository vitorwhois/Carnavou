package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/vitorwhois/Carnavou/routes"
	"github.com/vitorwhois/Carnavou/storage"
)

func main() {

	//Abre conexão com o banco de dados
	db, err := storage.OpenDatabaseConnection()
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados", err)
	}
	defer db.Close() //Fechar a conexão com o banco de dados

	routes.InitRoutes(db)

	port := 8080
	fmt.Printf("Servidor rodando em http://localhost:%d\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}

}
