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
	//Abre conexão com o banco de dados
	db, err := storage.OpenDatabaseConnection()
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados", err)
	}
	defer db.Close() //Fechar a conexão com o banco de dados

	// Crie uma instância do IndexHandler
	indexHandler := handlers.NewIndexHandler()
	handler := handlers.NewBlocosPorIDHandler(db)        // Passa a conexão do banco de dados
	minhalistaHandler := handlers.NewMinhalistaHandler() // Cria uma instância de MinhalistaHandler
	sobreHandler := handlers.NewSobreHandler()           // Cria uma instância de MinhalistaHandler

	// Use o método Handle da instância
	http.HandleFunc("/", indexHandler.Handle)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/blocosPorData", handlers.BlocosPorDataHandler)     // Ajuste da rota para /blocosPorData
	http.HandleFunc("/pesquisarBlocos", handlers.PesquisarBlocosHandler) // Pesquisa por nome
	http.Handle("/buscaid", handler)
	http.HandleFunc("/minhalista", minhalistaHandler.Handle)                //Lista criada pelo usuário
	http.HandleFunc("/sobre", sobreHandler.Handle)                          // Pagina Sobre nós
	http.HandleFunc("/subprefeitura", handlers.SubprefeituraBlocosHandler)  // Pesquisa por subprefeitura
	http.HandleFunc("/filtro", handlers.BlocosPorDataESubprefeituraHandler) // Pesquisa por data e subprefeitura

	port := 8080
	fmt.Printf("Servidor rodando em http://localhost:%d\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}

}
