package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/vitorwhois/Carnavou/handlers"
)

func main() {
	// Crie uma instância do IndexHandler
	indexHandler := handlers.NewIndexHandler()

	// Use o método Handle da instância
	http.HandleFunc("/", indexHandler.Handle)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := 8080
	fmt.Printf("Servidor rodando em http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}

}
