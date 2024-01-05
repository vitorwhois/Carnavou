package handlers

import (
	"fmt"
	"net/http"
)

// ListaHandler é o handler para a rota /lista
func ListaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Página de Lista - Em construção!")
}
