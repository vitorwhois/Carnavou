package handlers

import (
	"net/http"
	"text/template"

	"github.com/vitorwhois/Carnavou/models"
)

func ProgramacaoHandler(w http.ResponseWriter, r *http.Request) {
	repository := models.ListaRepository{}
	programacaoData, err := repository.ObterBlocosPorData("18/02/2024")
	if err != nil {
		http.Error(w, "Erro ao obter dados de programação", http.StatusInternalServerError)
		return
	}

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
