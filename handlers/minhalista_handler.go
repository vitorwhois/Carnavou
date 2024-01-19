package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/vitorwhois/Carnavou/models"
)

type BlocosPorIDHandler struct{}

func NewBlocosPorIDHandler() *BlocosPorIDHandler {
	return &BlocosPorIDHandler{}
}

func (h *BlocosPorIDHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling minhalista request")
	if r.Method == http.MethodGet {
		h.GetBlocosPorID(w, r)
	} else {
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
	}
}

func (h *BlocosPorIDHandler) GetBlocosPorID(w http.ResponseWriter, r *http.Request) {
	repository := models.ListaRepository{}

	ids := strings.Split(r.URL.Query().Get("ids"), ",")
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

	jsonData, err := json.Marshal(blocos)
	if err != nil {
		log.Printf("Erro ao converter dados para JSON: %v", err)
		http.Error(w, "Erro ao converter dados para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

/*package handlers

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
