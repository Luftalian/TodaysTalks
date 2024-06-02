package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/traPtitech/go-traq"
)

type Repository struct {
	db *sqlx.DB
}

type Repository2 struct {
	apiClient *traq.APIClient
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func New2(apiClient *traq.APIClient) *Repository2 {
	return &Repository2{apiClient: apiClient}
}
