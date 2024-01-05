package handlers

import (
	"net/http"
	"text/template"

	"github.com/vitorwhois/Carnavou/models"
)

// ProgramacaoHandler manipula solicitações para a página de programação
func ProgramacaoHandler(w http.ResponseWriter, r *http.Request) {
	// Lógica para obter dados de programação usando o serviço e repositório correspondentes
	repository := models.ListaRepository{}
	programacaoData, err := repository.ObterBlocosPorData("17/02/2024")
	if err != nil {
		http.Error(w, "Erro ao obter dados de programação", http.StatusInternalServerError)
		return
	}

	// Renderiza a página usando um template
	tmpl, err := template.ParseFiles("templates/programacao.html")
	if err != nil {
		http.Error(w, "Erro ao analisar o modelo", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, programacaoData)
	if err != nil {
		http.Error(w, "Erro ao executar o modelo", http.StatusInternalServerError)
	}
}
