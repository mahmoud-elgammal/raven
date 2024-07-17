package model

import (
	"github.com/mahmoud-elgammal/raven/pkg/database"
)

type UserRepository struct {
	*Repository
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func NewUser(db *database.DB) *UserRepository {
	return &UserRepository{
		Repository: NewRepository(db),
	}
}

func (repository *UserRepository) Create(user *UserRepository) {
	session := repository.db.GetSession()

	_ = session
}
