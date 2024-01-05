package handlers

import (
	"fmt"
	"net/http"
)

// NoticiasHandler é o handler para a rota /noticias
func NoticiasHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Página de Notícias - Em construção!")
}
