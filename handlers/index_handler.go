package handlers

import "net/http"

// IndexHandler representa o handler para a rota /
type IndexHandler struct{}

// NewIndexHandler cria uma nova instância de IndexHandler
func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

// Handle trata a requisição para a rota /
func (h *IndexHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Lógica específica para a página inicial
}
