package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/vitorwhois/Carnavou/models"
)

type BlocosPorIDHandler struct{}

func NewBlocosPorIDHandler() *BlocosPorIDHandler {
	return &BlocosPorIDHandler{}
}

func (h *BlocosPorIDHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling minhalista request")
	http.ServeFile(w, r, "templates/minhalista.html")
}

// GetBlocosPorID manipula solicitações para obter blocos por ID
func (h *BlocosPorIDHandler) GetBlocosPorID(w http.ResponseWriter, r *http.Request) {
	// Lógica para obter dados de blocos por ID usando o serviço e repositório correspondentes
	repository := models.ListaRepository{}

	// Recupera os IDs da solicitação HTTP atual
	ids := r.URL.Query()["ids"]
	if len(ids) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("IDs de blocos não fornecidos"))
		return
	}

	blocos, err := repository.ObterBlocosPorIDs(ids)
	if err != nil {
		log.Printf("Erro ao obter dados de blocos por ID: %v", err)
		http.Error(w, "Erro ao obter dados de blocos por ID", http.StatusInternalServerError)
		return
	}

	// Gera o HTML para cada bloco
	html := "<div id='minhalista'>"
	for _, bloco := range blocos {
		cardHtml := criarCard(bloco) // Chama a função criarCard para cada bloco
		html += cardHtml
	}
	html += "</div>"

	// Define o cabeçalho Content-Type para text/html
	w.Header().Set("Content-Type", "text/html")

	// Responde com o HTML gerado
	w.Write([]byte(html))
}

func criarCard(bloco models.Bloco) string {
	// Cria o elemento card
	card := `<div class="card mb-2">`

	// Adiciona a div com o título
	cardTitulo := `
<div class="card-titulo" style="display: flex; justify-content: space-between;">
    <h3>` + bloco.Nome + `</h3>
    <button class="add-list-link" data-bloco-id="` + strconv.Itoa(bloco.ID) + `">Add a lista <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-plus" viewBox="0 0 16 16">
        <path d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4"></path>
    </svg></button>
</div>`
	card += cardTitulo

	// Adiciona a div com as informações do card
	infoCard := `<div class="info-card">`

	// Adiciona a div com a data
	cardData := `
    <div class="card-data">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-calendar" viewBox="0 0 16 16">
            <path d="M3.5 0a.5.5 0 0 1 .5.5V1h8V.5a.5.5 0 0 1 1 0V1h1a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2h1V.5a.5.5 0 0 1 .5-.5M1 4v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V4z"></path>
        </svg>
        <span>` + bloco.Data[0:5] + `</span>
    </div>`
	infoCard += cardData

	// Adiciona a div com o endereço
	cardEndereco := `
    <div class="card-endereco">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-geo-alt" viewBox="0 0 16 16">
            <path d="M12.166 8.94c-.524 1.062-1.234 2.12-1.96 3.07A32 32 0 0 1 8 14.58a32 32 0 0 1-2.206-2.57c-.726-.95-1.436-2.008-1.96-3.07C3.304 7.867 3 6.862 3 6a5 5 0 0 1 10 0c0 .862-.305 1.867-.834 2.94M8 16s6-5.686 6-10A6 6 0 0 0 2 6c0 4.314 6 10 6 10"/>
            <path d="M8 8a2 2 0 1 1 0-4 2 2 0 0 1 0 4m0 1a3 3 0 1 0 0-6 3 3 0 0 0 0 6"/>
        </svg>
        <span>` + bloco.Local + `</span>
    </div>`
	infoCard += cardEndereco

	// Adiciona a div com o horário
	cardHorario := `
    <div class="card-horario">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-clock" viewBox="0 0 16 16">
            <path d="M8 3.5a.5.5 0 0 0-1 0V9a.5.5 0 0 0 .252.434l3.5 2a.5.5 0 0 0 .496-.868L8 8.71z"></path>
            <path d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16m7-8A7 7 0 1 1 1 8a7 7 0 0 1 14 0"></path>
        </svg>
        <span>` + bloco.Concentracao[0:2] + `h</span>
    </div>`
	infoCard += cardHorario

	// Adiciona a div de informações ao card
	card += infoCard

	// Fecha a div do card
	card += `</div>`

	return card
}

/*
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vitorwhois/Carnavou/models"
)

type BlocosPorIDHandler struct{}

func NewBlocosPorIDHandler() *BlocosPorIDHandler {
	return &BlocosPorIDHandler{}
}

func (h *BlocosPorIDHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling minhalista request")
	http.ServeFile(w, r, "templates/minhalista.html")
}

// GetBlocosPorID manipula solicitações para obter blocos por ID
func (h *BlocosPorIDHandler) GetBlocosPorID(w http.ResponseWriter, r *http.Request) {
	// Lógica para obter dados de blocos por ID usando o serviço e repositório correspondentes
	repository := models.ListaRepository{}

	// Recupera os IDs da solicitação HTTP atual
	ids := r.URL.Query()["ids"]
	if len(ids) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("IDs de blocos não fornecidos"))
		return
	}

	blocos, err := repository.ObterBlocosPorIDs(ids)
	if err != nil {
		log.Printf("Erro ao obter dados de blocos por ID: %v", err)
		http.Error(w, "Erro ao obter dados de blocos por ID", http.StatusInternalServerError)
		return
	}

	// Converte os dados para JSON
	jsonData, err := json.Marshal(blocos)
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
*/
