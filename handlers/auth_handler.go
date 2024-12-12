package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/vitorwhois/Carnavou/repositories"
)

type AuthHandler struct {
	repo *repositories.UserRepository
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	repo := repositories.NewUserRepository(db)
	return &AuthHandler{repo: repo}
}

// Handle para cadastro de usuário
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	if err := h.repo.CreateUser(user.Email, user.Password); err != nil {
		log.Printf("Erro ao cadastrar usuário: %v", err)
		http.Error(w, "Erro ao cadastrar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Handle para login de usuário
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	if valid, err := h.repo.ValidateUser(user.Email, user.Password); err != nil || !valid {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "Credenciais inválidas"})
		return
	}

	// Aqui você pode gerar um token JWT ou outra forma de autenticação, se necessário
	// token := generateToken(user.Email) // Exemplo de geração de token

	// Resposta de sucesso
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login bem-sucedido",
		// "token": token, // Inclua o token se estiver usando autenticação baseada em token
	})
} 