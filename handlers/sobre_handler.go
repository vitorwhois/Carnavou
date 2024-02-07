package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
)

type SobreHandler struct{}

func NewSobreHandler() *SobreHandler {
	return &SobreHandler{}
}

func (h *SobreHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling sobre nós request")
	// Determine the absolute path of the project
	absPath, _ := filepath.Abs(".")
	// Build the path to the file from the absolute path of the project
	filePath := filepath.Join(absPath, "templates", "sobre.html")
	http.ServeFile(w, r, filePath)
}
