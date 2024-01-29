package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
)

type MinhalistaHandler struct{}

func NewMinhalistaHandler() *MinhalistaHandler {
	return &MinhalistaHandler{}
}

func (h *MinhalistaHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling minhalista request")
	// Determine the absolute path of the project
	absPath, _ := filepath.Abs(".")
	// Build the path to the file from the absolute path of the project
	filePath := filepath.Join(absPath, "templates", "minhalista.html")
	http.ServeFile(w, r, filePath)
}
