package model

import (
	"github.com/mahmoud-elgammal/raven/pkg/database"
)

// Repository manages interactions with Neo4j database.
type Repository struct {
	db *database.DB
}

type RepositoryInterface interface {
	Create(repository RepositoryInterface)
}

func NewRepository(db *database.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (*Repository) Create(repository RepositoryInterface) {
}
