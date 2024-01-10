package handlers

import (
	"fmt"
	"net/http"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling index request")
	http.ServeFile(w, r, "templates/index.html")
}
