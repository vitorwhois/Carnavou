package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	// SupabaseDB está disponível para ser usado
	//_ = supabase.SupabaseDB

	// Configuração de rotas
	http.HandleFunc("/", IndexHandler)
	// ... outras rotas

	// Inicia o servidor local na porta 8080
	port := 8080
	fmt.Printf("Servidor rodando em http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}

// IndexHandler é o manipulador para a página index.html
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
