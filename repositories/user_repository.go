package repositories

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, hashedPassword)
	return err
}

func (r *UserRepository) ValidateUser(email, password string) (bool, error) {
	var hashedPassword string
	err := r.db.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil, nil
} 